package parser

import (
	"fmt"
	"sort"
	"strings"

	"github.com/go-openapi/spec"
	"github.com/hashicorp/pandora/tools/importer-rest-api-specs/cleanup"
	"github.com/hashicorp/pandora/tools/importer-rest-api-specs/models"
)

type resourceIdParseResult struct {
	// nameToResourceIDs is a map[name]ParsedResourceID containing information about the Resource ID's
	nameToResourceIDs map[string]models.ParsedResourceId

	// nestedResult contains any constants and models found when parsing these ID's
	nestedResult parseResult

	// resourceUriMetadata is a map[uri]resourceUriMetadata for the Resource ID reference & any suffix
	resourceUrisToMetadata map[string]resourceUriMetadata
}

type resourceUriMetadata struct {
	// resourceIdName is the name of the ResourceID object, available once the unique names have been
	// identified (if there's a Resource ID)
	resourceIdName *string

	// resourceId is the parsed Resource ID object, if any
	resourceId *models.ParsedResourceId

	// uriSuffix is any suffix which should be applied to the URI
	uriSuffix *string
}

func (d *SwaggerDefinition) findResourceIdsForTag(tag *string) (*resourceIdParseResult, error) {
	result := resourceIdParseResult{
		nestedResult: parseResult{
			constants: map[string]models.ConstantDetails{},
		},

		nameToResourceIDs:      map[string]models.ParsedResourceId{},
		resourceUrisToMetadata: map[string]resourceUriMetadata{},
	}

	// first get a list of all of the Resource ID's present in these operations
	// where a Suffix is present on a Resource ID, we'll have 2 entries for the Suffix and the Resource ID directly
	urisToMetadata, nestedResult, err := d.parseResourceIdsFromOperations(tag)
	if err != nil {
		return nil, fmt.Errorf("parsing Resource ID's from Operations: %+v", err)
	}
	result.nestedResult = *nestedResult

	// next determine names for these
	namesToResourceUris, urisToNames, err := determineNamesForResourceIds(*urisToMetadata)
	if err != nil {
		return nil, fmt.Errorf("determining names for Resource ID's: %+v", err)
	}
	result.nameToResourceIDs = *namesToResourceUris

	// finally go over the existing results and swap out the Resource ID objects for the Name which should be used
	urisToMetadata, err = mapNamesToResourceIds(*urisToNames, *urisToMetadata)
	if err != nil {
		return nil, fmt.Errorf("mapping names back to Resource ID's: %+v", err)
	}
	result.resourceUrisToMetadata = *urisToMetadata

	return &result, nil
}

func (d SwaggerDefinition) parseResourceIdsFromOperations(tag *string) (*map[string]resourceUriMetadata, *parseResult, error) {
	result := parseResult{
		constants: map[string]models.ConstantDetails{},
	}
	urisToMetaData := make(map[string]resourceUriMetadata, 0)

	for _, operation := range d.swaggerSpecExpanded.Operations() {
		for uri, operationDetails := range operation {
			if !operationMatchesTag(operationDetails, tag) {
				continue
			}

			if operationShouldBeIgnored(uri) {
				continue
			}

			metadata, err := d.parseResourceIdFromOperation(uri, operationDetails)
			if err != nil {
				return nil, nil, fmt.Errorf("parsing %q: %+v", uri, err)
			}

			// next, if it's based on a Resource ID, let's ensure that's added too
			resourceUri := uri
			if metadata.resourceId != nil {
				result.appendConstants(metadata.resourceId.Constants)

				resourceManagerUri := metadata.resourceId.NormalizedResourceManagerResourceId()
				if resourceUri != resourceManagerUri {
					urisToMetaData[resourceManagerUri] = resourceUriMetadata{
						resourceIdName: metadata.resourceIdName,
						resourceId:     metadata.resourceId,
						uriSuffix:      nil,
					}
				}
			}
			urisToMetaData[resourceUri] = *metadata
		}
	}

	return &urisToMetaData, &result, nil
}

func (d *SwaggerDefinition) parseResourceIdFromOperation(uri string, operationDetails *spec.Operation) (*resourceUriMetadata, error) {
	// TODO: unit tests for this method too
	segments := make([]models.ResourceIdSegment, 0)
	result := parseResult{
		constants: map[string]models.ConstantDetails{},
	}

	uriSegments := strings.Split(strings.TrimPrefix(uri, "/"), "/")
	for _, uriSegment := range uriSegments {
		originalSegment := uriSegment
		normalizedSegment := cleanup.RemoveInvalidCharacters(uriSegment, false)
		normalizedSegment = cleanup.NormalizeSegment(normalizedSegment, true)
		// the names should always be camelCased, so let's be sure
		normalizedSegment = fmt.Sprintf("%s%s", strings.ToLower(string(normalizedSegment[0])), normalizedSegment[1:])

		// intentionally check the pre-cut version
		if strings.HasPrefix(originalSegment, "{") && strings.HasSuffix(originalSegment, "}") {
			if strings.EqualFold(normalizedSegment, "scope") {
				segments = append(segments, models.ResourceIdSegment{
					Type: models.ScopeSegment,
					Name: normalizedSegment,
				})
				continue
			}

			if strings.EqualFold(normalizedSegment, "subscriptionId") {
				segments = append(segments, models.ResourceIdSegment{
					Type: models.SubscriptionIdSegment,
					Name: normalizedSegment,
				})
				continue
			}

			if strings.EqualFold(normalizedSegment, "resourceGroupName") {
				segments = append(segments, models.ResourceIdSegment{
					Type: models.ResourceGroupSegment,
					Name: normalizedSegment,
				})
				continue
			}

			isConstant := false
			for _, param := range operationDetails.Parameters {
				if strings.EqualFold(param.Name, normalizedSegment) && strings.EqualFold(param.In, "path") {
					if param.Ref.String() != "" {
						return nil, fmt.Errorf("TODO: Enum's aren't supported by Reference right now, but apparently should be for %q", uriSegment)
					}

					if param.Enum != nil {
						// then find the constant itself
						constant, err := mapConstant([]string{param.Type}, param.Name, param.Enum, param.Extensions)
						if err != nil {
							return nil, fmt.Errorf("parsing constant from %q: %+v", uriSegment, err)
						}
						result.constants[constant.name] = constant.details
						segments = append(segments, models.ResourceIdSegment{
							Type:              models.ConstantSegment,
							Name:              normalizedSegment,
							ConstantReference: &constant.name,
						})
						isConstant = true
						break
					}
				}
			}
			if isConstant {
				continue
			}

			segments = append(segments, models.ResourceIdSegment{
				Type: models.UserSpecifiedSegment,
				Name: normalizedSegment,
			})
			continue
		}

		segments = append(segments, models.ResourceIdSegment{
			Type:       models.StaticSegment,
			Name:       normalizedSegment,
			FixedValue: &originalSegment,
		})
	}

	output := resourceUriMetadata{
		resourceIdName: nil,
		uriSuffix:      nil,
	}

	// UriSuffixes are "operations" on a given Resource ID/URI - for example `/restart`
	// or in the case of List operations /providers/Microsoft.Blah/listAllTheThings
	// we treat these as "operations" on the Resource ID and as such the "segments" should
	// only be for the Resource ID and not for the UriSuffix (which is as an additional field)
	lastUserValueSegment := -1
	for i, segment := range segments {
		// everything else technically is a user configurable component
		if segment.Type != models.StaticSegment {
			lastUserValueSegment = i
		}
	}
	if lastUserValueSegment >= 0 && len(segments) > lastUserValueSegment+1 {
		suffix := ""
		for _, segment := range segments[lastUserValueSegment+1:] {
			suffix += fmt.Sprintf("/%s", *segment.FixedValue)
		}
		output.uriSuffix = &suffix

		// remove any URI Suffix since this isn't relevant for the ID's
		segments = segments[0 : lastUserValueSegment+1]
	}

	allSegmentsAreStatic := true
	for _, segment := range segments {
		if segment.Type != models.StaticSegment {
			allSegmentsAreStatic = false
			break
		}
	}
	if allSegmentsAreStatic {
		// if it's not an ARM ID there's nothing to output here, but new up a placeholder
		// to be able to give us a normalized id for the suffix
		pri := models.ParsedResourceId{
			Constants: result.constants,
			Segments:  segments,
		}
		suffix := pri.NormalizedResourceId()
		output.uriSuffix = &suffix
	} else {
		output.resourceId = &models.ParsedResourceId{
			Constants: result.constants,
			Segments:  segments,
		}
	}

	return &output, nil
}

// determineNamesForResourceIds returns a map[name]ParsedResourceID and map[Uri]Name based on the Resource Manager URI's available
func determineNamesForResourceIds(urisToObjects map[string]resourceUriMetadata) (*map[string]models.ParsedResourceId, *map[string]string, error) {
	// now that we have all of the Resource ID's, we then need to go through and determine Unique ID's for those
	// we need all of them here to avoid conflicts, e.g. AuthorizationRule which can be a NamespaceAuthorizationRule
	// or an EventHubAuthorizationRule, but is named AuthorizationRule in both

	// Before we do anything else, let's go through remove any containing uri suffixes (since these are duplicated without
	// where they contain a Resource ID - and then sort them short -> long for consistency
	sortedUris := make([]string, 0)
	for uri, resourceId := range urisToObjects {
		// if it's just a suffix (e.g. root-level ListAll calls) iterate over it
		if resourceId.resourceId == nil {
			continue
		}

		// when there's a Uri Suffix we should pass in both the full uri and just the resource manager uri so we can
		// skip it if this is a full uri (with a suffix), since the name comes from the resource manager uri instead
		if resourceId.uriSuffix != nil {
			continue
		}

		sortedUris = append(sortedUris, uri)
	}

	// sort these by length
	sort.Slice(sortedUris, func(x, y int) bool {
		return len(sortedUris[x]) < len(sortedUris[y])
	})

	candidateNamesToUris := make(map[string]models.ParsedResourceId, 0)
	conflictingNamesToUris := make(map[string][]models.ParsedResourceId, 0)
	for _, uri := range sortedUris {
		resourceId := urisToObjects[uri]

		// NOTE: these are returned sorted from right to left in URI's, since they're assumed to be hierarchical
		segmentsAvailableForNaming := resourceId.resourceId.SegmentsAvailableForNaming()
		if len(segmentsAvailableForNaming) == 0 {
			return nil, nil, fmt.Errorf("the uri %q has no segments available for naming", segmentsAvailableForNaming)
		}

		candidateSegmentName := segmentsAvailableForNaming[0]
		if resourceId.resourceId.Segments[0].Type == models.ScopeSegment && len(resourceId.resourceId.Segments) > 1 {
			candidateSegmentName = fmt.Sprintf("Scoped%s", candidateSegmentName)
		}
		candidateSegmentName = cleanup.NormalizeSegment(candidateSegmentName, false)

		// if we have an existing conflicting key, let's add this to that
		if uris, existing := conflictingNamesToUris[candidateSegmentName]; existing {
			uris = append(uris, *resourceId.resourceId)
			conflictingNamesToUris[candidateSegmentName] = uris
			continue
		}

		// if there's an existing candidate name for this key, move both this URI and that one to the Conflicts
		if existingUri, existing := candidateNamesToUris[candidateSegmentName]; existing {
			conflictingNamesToUris[candidateSegmentName] = []models.ParsedResourceId{existingUri, *resourceId.resourceId}
			delete(candidateNamesToUris, candidateSegmentName)
			continue
		}

		// otherwise we have a candidate name we should be able to use, so let's run with it
		candidateNamesToUris[candidateSegmentName] = *resourceId.resourceId
	}

	// now we need to fix the conflicts
	for _, conflictingUris := range conflictingNamesToUris {
		uniqueNames, err := determineUniqueNamesFor(conflictingUris, candidateNamesToUris)
		if err != nil {
			uris := make([]string, 0)
			for _, uri := range conflictingUris {
				uris = append(uris, uri.String())
			}

			return nil, nil, fmt.Errorf("determining unique names for conflicting uri's %q: %+v", strings.Join(uris, " | "), err)
		}

		for k, v := range *uniqueNames {
			candidateNamesToUris[k] = v
		}
	}

	// now we have unique ID's, we should go through and suffix `Id` onto the end of each of them
	outputNamesToUris := make(map[string]models.ParsedResourceId)
	for k, v := range candidateNamesToUris {
		key := fmt.Sprintf("%sId", k)
		outputNamesToUris[key] = v
	}

	// finally compose a list of uris -> names so these are easier to map back
	urisToNames := make(map[string]string, 0)
	for k, v := range outputNamesToUris {
		urisToNames[v.NormalizedResourceManagerResourceId()] = k
	}

	return &outputNamesToUris, &urisToNames, nil
}

func determineUniqueNamesFor(conflictingUris []models.ParsedResourceId, existingCandidateNames map[string]models.ParsedResourceId) (*map[string]models.ParsedResourceId, error) {
	proposedNames := make(map[string]models.ParsedResourceId)
	for _, resourceId := range conflictingUris {
		availableSegments := resourceId.SegmentsAvailableForNaming()

		proposedName := ""
		uniqueNameFound := false

		// matches the behaviour above
		if resourceId.Segments[0].Type == models.ScopeSegment {
			proposedName += "Scoped"
		}

		for _, segment := range availableSegments {
			proposedName = fmt.Sprintf("%s%s", cleanup.NormalizeSegment(segment, false), proposedName)

			_, hasConflictWithExisting := existingCandidateNames[proposedName]
			_, hasConflictWithProposed := proposedNames[proposedName]
			if !hasConflictWithProposed && !hasConflictWithExisting {
				uniqueNameFound = true
				break
			}
		}

		if !uniqueNameFound {
			return nil, fmt.Errorf("not enough segments in %q to determine a unique name", resourceId.String())
		}

		proposedNames[proposedName] = resourceId
	}

	return &proposedNames, nil
}

func mapNamesToResourceIds(urisToNames map[string]string, urisToMetadata map[string]resourceUriMetadata) (*map[string]resourceUriMetadata, error) {
	output := make(map[string]resourceUriMetadata, 0)

	for uri, metadata := range urisToMetadata {
		// ID's with just Suffixes are valid and won't have an ID Type, so skip those
		if metadata.resourceId == nil {
			output[uri] = metadata
			continue
		}

		name, ok := urisToNames[metadata.resourceId.NormalizedResourceManagerResourceId()]
		if !ok {
			return nil, fmt.Errorf("Resource ID : Name mapping not found for %q", uri)
		}

		output[metadata.resourceId.NormalizedResourceManagerResourceId()] = resourceUriMetadata{
			resourceIdName: &name,
			// intentionally don't map over the UriSuffix since this is handled above
		}

		// when there's a suffix, we need to output the full uri in the map too
		if metadata.uriSuffix != nil {
			metadata.resourceIdName = &name
			output[uri] = metadata
		}
	}

	return &output, nil
}