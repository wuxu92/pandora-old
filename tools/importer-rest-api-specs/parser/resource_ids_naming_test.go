package parser

import (
	"reflect"
	"testing"

	"github.com/hashicorp/pandora/tools/importer-rest-api-specs/models"
)

var subscriptionResourceId = models.ParsedResourceId{
	Constants: map[string]models.ConstantDetails{},
	Segments: []models.ResourceIdSegment{
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("subscriptions"),
			Name:       "subscriptions",
		},
		{
			Name: "subscriptionId",
			Type: models.SubscriptionIdSegment,
		},
	},
}
var resourceGroupResourceId = models.ParsedResourceId{
	Constants: map[string]models.ConstantDetails{},
	Segments: []models.ResourceIdSegment{
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("subscriptions"),
			Name:       "subscriptions",
		},
		{
			Name: "subscriptionId",
			Type: models.SubscriptionIdSegment,
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("resourceGroups"),
			Name:       "resourceGroups",
		},
		{
			Name: "resourceGroupName",
			Type: models.ResourceGroupSegment,
		},
	},
}
var managementGroupResourceId = models.ParsedResourceId{
	Constants: map[string]models.ConstantDetails{},
	Segments: []models.ResourceIdSegment{
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("providers"),
			Name:       "providers",
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("Microsoft.Management"),
			Name:       "Microsoft.Management",
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("managementGroups"),
			Name:       "managementGroups",
		},
		{
			Name: "name",
			Type: models.UserSpecifiedSegment,
		},
	},
}
var virtualMachineResourceId = models.ParsedResourceId{
	Constants: map[string]models.ConstantDetails{},
	Segments: []models.ResourceIdSegment{
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("subscriptions"),
			Name:       "subscriptions",
		},
		{
			Name: "subscriptionId",
			Type: models.SubscriptionIdSegment,
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("resourceGroups"),
			Name:       "resourceGroups",
		},
		{
			Name: "resourceGroupName",
			Type: models.ResourceGroupSegment,
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("providers"),
			Name:       "providers",
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("Microsoft.Compute"),
			Name:       "Microsoft.Compute",
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("virtualMachines"),
			Name:       "virtualMachines",
		},
		{
			Name: "virtualMachineName",
			Type: models.UserSpecifiedSegment,
		},
	},
}
var virtualMachineExtensionResourceId = models.ParsedResourceId{
	Constants: map[string]models.ConstantDetails{},
	Segments: []models.ResourceIdSegment{
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("subscriptions"),
			Name:       "subscriptions",
		},
		{
			Name: "subscriptionId",
			Type: models.SubscriptionIdSegment,
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("resourceGroups"),
			Name:       "resourceGroups",
		},
		{
			Name: "resourceGroupName",
			Type: models.ResourceGroupSegment,
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("providers"),
			Name:       "providers",
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("Microsoft.Compute"),
			Name:       "Microsoft.Compute",
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("virtualMachines"),
			Name:       "virtualMachines",
		},
		{
			Name: "virtualMachineName",
			Type: models.UserSpecifiedSegment,
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("extensions"),
			Name:       "extensions",
		},
		{
			Name: "extensionName",
			Type: models.UserSpecifiedSegment,
		},
	},
}
var virtualNetworkExtensionResourceId = models.ParsedResourceId{
	Constants: map[string]models.ConstantDetails{},
	Segments: []models.ResourceIdSegment{
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("subscriptions"),
			Name:       "subscriptions",
		},
		{
			Name: "subscriptionId",
			Type: models.SubscriptionIdSegment,
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("resourceGroups"),
			Name:       "resourceGroups",
		},
		{
			Name: "resourceGroupName",
			Type: models.ResourceGroupSegment,
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("providers"),
			Name:       "providers",
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("Microsoft.Network"),
			Name:       "Microsoft.Network",
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("extensions"),
			Name:       "extensions",
		},
		{
			Name: "extensionName",
			Type: models.UserSpecifiedSegment,
		},
	},
}
var scopedMonitorResourceId = models.ParsedResourceId{
	Constants: map[string]models.ConstantDetails{},
	Segments: []models.ResourceIdSegment{
		{
			Type: models.ScopeSegment,
			Name: "scope",
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("providers"),
			Name:       "providers",
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("Microsoft.Monitor"),
			Name:       "Microsoft.Monitor",
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("extensions"),
			Name:       "extensions",
		},
		{
			Name: "extensionName",
			Type: models.UserSpecifiedSegment,
		},
	},
}
var signalRResourceId = models.ParsedResourceId{
	Constants: map[string]models.ConstantDetails{},
	Segments: []models.ResourceIdSegment{
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("subscriptions"),
			Name:       "subscriptions",
		},
		{
			Name: "subscriptionId",
			Type: models.SubscriptionIdSegment,
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("resourceGroups"),
			Name:       "resourceGroups",
		},
		{
			Name: "resourceGroupName",
			Type: models.ResourceGroupSegment,
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("providers"),
			Name:       "providers",
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("Microsoft.SignalRService"),
			Name:       "Microsoft.SignalRService",
		},
		{
			Type:       models.StaticSegment,
			FixedValue: strPtr("SignalR"),
			Name:       "SignalR",
		},
		{
			Name: "resourceName",
			Type: models.UserSpecifiedSegment,
		},
	},
}

func TestResourceIDNamingEmpty(t *testing.T) {
	actualNamesToIds, actualUrisToNames, err := determineNamesForResourceIds(map[string]resourceUriMetadata{})
	if err != nil {
		t.Fatalf("error: %+v", err)
		return
	}

	if len(*actualNamesToIds) > 0 {
		t.Fatalf("expected there to be no namesToIds but got %+v", *actualNamesToIds)
	}

	if len(*actualUrisToNames) > 0 {
		t.Fatalf("expected there to be no urisToNames but got %+v", *actualUrisToNames)
	}
}

func TestResourceIDNamingSubscriptionId(t *testing.T) {
	input := map[string]resourceUriMetadata{
		"/subscriptions/{subscriptionId}": {
			resourceIdName: nil,
			resourceId:     &subscriptionResourceId,
			uriSuffix:      nil,
		},
	}
	expectedNamesToIds := map[string]models.ParsedResourceId{
		"SubscriptionId": subscriptionResourceId,
	}
	expectedUrisToNames := map[string]string{
		"/subscriptions/{subscriptionId}": "SubscriptionId",
	}

	actualNamesToIds, actualUrisToNames, err := determineNamesForResourceIds(input)
	if err != nil {
		t.Fatalf("error: %+v", err)
		return
	}

	if !reflect.DeepEqual(expectedNamesToIds, *actualNamesToIds) {
		t.Fatalf("expected namesToIds to be %+v but got %+v", expectedNamesToIds, *actualNamesToIds)
	}

	if !reflect.DeepEqual(expectedUrisToNames, *actualUrisToNames) {
		t.Fatalf("expected urisToNames to be %+v but got %+v", expectedUrisToNames, *actualUrisToNames)
	}
}

func TestResourceIDNamingSubscriptionIdAndSuffix(t *testing.T) {
	input := map[string]resourceUriMetadata{
		"/subscriptions/{subscriptionId}": {
			resourceIdName: nil,
			resourceId:     &subscriptionResourceId,
			uriSuffix:      nil,
		},
		"/subscriptions/{subscriptionId}/resourceGroups": {
			resourceIdName: nil,
			resourceId:     &subscriptionResourceId,
			uriSuffix:      strPtr("/resourceGroups"),
		},
	}
	expectedNamesToIds := map[string]models.ParsedResourceId{
		"SubscriptionId": subscriptionResourceId,
	}
	expectedUrisToNames := map[string]string{
		"/subscriptions/{subscriptionId}": "SubscriptionId",
	}

	actualNamesToIds, actualUrisToNames, err := determineNamesForResourceIds(input)
	if err != nil {
		t.Fatalf("error: %+v", err)
		return
	}

	if !reflect.DeepEqual(expectedNamesToIds, *actualNamesToIds) {
		t.Fatalf("expected namesToIds to be %+v but got %+v", expectedNamesToIds, *actualNamesToIds)
	}

	if !reflect.DeepEqual(expectedUrisToNames, *actualUrisToNames) {
		t.Fatalf("expected urisToNames to be %+v but got %+v", expectedUrisToNames, *actualUrisToNames)
	}
}

func TestResourceIDNamingResourceGroupId(t *testing.T) {
	input := map[string]resourceUriMetadata{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}": {
			resourceIdName: nil,
			resourceId:     &resourceGroupResourceId,
			uriSuffix:      nil,
		},
	}
	expectedNamesToIds := map[string]models.ParsedResourceId{
		"ResourceGroupId": resourceGroupResourceId,
	}
	expectedUrisToNames := map[string]string{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}": "ResourceGroupId",
	}

	actualNamesToIds, actualUrisToNames, err := determineNamesForResourceIds(input)
	if err != nil {
		t.Fatalf("error: %+v", err)
		return
	}

	if !reflect.DeepEqual(expectedNamesToIds, *actualNamesToIds) {
		t.Fatalf("expected namesToIds to be %+v but got %+v", expectedNamesToIds, *actualNamesToIds)
	}

	if !reflect.DeepEqual(expectedUrisToNames, *actualUrisToNames) {
		t.Fatalf("expected urisToNames to be %+v but got %+v", expectedUrisToNames, *actualUrisToNames)
	}
}

func TestResourceIDNamingResourceGroupIdAndSuffix(t *testing.T) {
	input := map[string]resourceUriMetadata{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}": {
			resourceIdName: nil,
			resourceId:     &resourceGroupResourceId,
			uriSuffix:      nil,
		},
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/restart": {
			resourceIdName: nil,
			resourceId:     &resourceGroupResourceId,
			uriSuffix:      strPtr("/restart"),
		},
	}
	expectedNamesToIds := map[string]models.ParsedResourceId{
		"ResourceGroupId": resourceGroupResourceId,
	}
	expectedUrisToNames := map[string]string{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}": "ResourceGroupId",
	}

	actualNamesToIds, actualUrisToNames, err := determineNamesForResourceIds(input)
	if err != nil {
		t.Fatalf("error: %+v", err)
		return
	}

	if !reflect.DeepEqual(expectedNamesToIds, *actualNamesToIds) {
		t.Fatalf("expected namesToIds to be %+v but got %+v", expectedNamesToIds, *actualNamesToIds)
	}

	if !reflect.DeepEqual(expectedUrisToNames, *actualUrisToNames) {
		t.Fatalf("expected urisToNames to be %+v but got %+v", expectedUrisToNames, *actualUrisToNames)
	}
}

func TestResourceIDNamingManagementGroupId(t *testing.T) {
	input := map[string]resourceUriMetadata{
		"/providers/Microsoft.Management/managementGroups/{name}": {
			resourceIdName: nil,
			resourceId:     &managementGroupResourceId,
			uriSuffix:      nil,
		},
	}
	expectedNamesToIds := map[string]models.ParsedResourceId{
		"ManagementGroupId": managementGroupResourceId,
	}
	expectedUrisToNames := map[string]string{
		"/providers/Microsoft.Management/managementGroups/{name}": "ManagementGroupId",
	}

	actualNamesToIds, actualUrisToNames, err := determineNamesForResourceIds(input)
	if err != nil {
		t.Fatalf("error: %+v", err)
		return
	}

	if !reflect.DeepEqual(expectedNamesToIds, *actualNamesToIds) {
		t.Fatalf("expected namesToIds to be %+v but got %+v", expectedNamesToIds, *actualNamesToIds)
	}

	if !reflect.DeepEqual(expectedUrisToNames, *actualUrisToNames) {
		t.Fatalf("expected urisToNames to be %+v but got %+v", expectedUrisToNames, *actualUrisToNames)
	}
}

func TestResourceIDNamingManagementGroupIdAndSuffix(t *testing.T) {
	input := map[string]resourceUriMetadata{
		"/providers/Microsoft.Management/managementGroups/{name}": {
			resourceIdName: nil,
			resourceId:     &managementGroupResourceId,
			uriSuffix:      nil,
		},
		"/providers/Microsoft.Management/managementGroups/{name}/restart": {
			resourceIdName: nil,
			resourceId:     &managementGroupResourceId,
			uriSuffix:      strPtr("/restart"),
		},
	}
	expectedNamesToIds := map[string]models.ParsedResourceId{
		"ManagementGroupId": managementGroupResourceId,
	}
	expectedUrisToNames := map[string]string{
		"/providers/Microsoft.Management/managementGroups/{name}": "ManagementGroupId",
	}

	actualNamesToIds, actualUrisToNames, err := determineNamesForResourceIds(input)
	if err != nil {
		t.Fatalf("error: %+v", err)
		return
	}

	if !reflect.DeepEqual(expectedNamesToIds, *actualNamesToIds) {
		t.Fatalf("expected namesToIds to be %+v but got %+v", expectedNamesToIds, *actualNamesToIds)
	}

	if !reflect.DeepEqual(expectedUrisToNames, *actualUrisToNames) {
		t.Fatalf("expected urisToNames to be %+v but got %+v", expectedUrisToNames, *actualUrisToNames)
	}
}

func TestResourceIDNamingTopLevelScope(t *testing.T) {
	scopeResourceId := models.ParsedResourceId{
		Constants: map[string]models.ConstantDetails{},
		Segments: []models.ResourceIdSegment{
			{
				Type: models.ScopeSegment,
				Name: "scope",
			},
		},
	}

	input := map[string]resourceUriMetadata{
		"/{scope}": {
			resourceIdName: nil,
			resourceId:     &scopeResourceId,
			uriSuffix:      nil,
		},
	}
	expectedNamesToIds := map[string]models.ParsedResourceId{
		"ScopeId": scopeResourceId,
	}
	expectedUrisToNames := map[string]string{
		"/{scope}": "ScopeId",
	}

	actualNamesToIds, actualUrisToNames, err := determineNamesForResourceIds(input)
	if err != nil {
		t.Fatalf("error: %+v", err)
		return
	}

	if !reflect.DeepEqual(expectedNamesToIds, *actualNamesToIds) {
		t.Fatalf("expected namesToIds to be %+v but got %+v", expectedNamesToIds, *actualNamesToIds)
	}

	if !reflect.DeepEqual(expectedUrisToNames, *actualUrisToNames) {
		t.Fatalf("expected urisToNames to be %+v but got %+v", expectedUrisToNames, *actualUrisToNames)
	}
}

func TestResourceIDNamingContainingAConstant(t *testing.T) {
	dnsResourceId := models.ParsedResourceId{
		Constants: map[string]models.ConstantDetails{
			"DnsRecordType": {
				FieldType: models.StringConstant,
				Values: map[string]string{
					"A":    "A",
					"AAAA": "AAAA",
				},
			},
		},
		Segments: []models.ResourceIdSegment{
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("subscriptions"),
				Name:       "subscriptions",
			},
			{
				Name: "subscriptionId",
				Type: models.SubscriptionIdSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("resourceGroups"),
				Name:       "resourceGroups",
			},
			{
				Name: "resourceGroupName",
				Type: models.ResourceGroupSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("providers"),
				Name:       "providers",
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("Microsoft.Network"),
				Name:       "Microsoft.Network",
			},
			{
				Type:              models.ConstantSegment,
				Name:              "recordType",
				ConstantReference: strPtr("DnsRecordType"),
			},
			{
				Name: "recordName",
				Type: models.UserSpecifiedSegment,
			},
		},
	}

	input := map[string]resourceUriMetadata{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/{recordType}/{recordName}": {
			resourceIdName: nil,
			resourceId:     &dnsResourceId,
			uriSuffix:      nil,
		},
	}
	expectedNamesToIds := map[string]models.ParsedResourceId{
		"RecordTypeId": dnsResourceId,
	}
	expectedUrisToNames := map[string]string{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/{recordType}/{recordName}": "RecordTypeId",
	}

	actualNamesToIds, actualUrisToNames, err := determineNamesForResourceIds(input)
	if err != nil {
		t.Fatalf("error: %+v", err)
		return
	}

	if !reflect.DeepEqual(expectedNamesToIds, *actualNamesToIds) {
		t.Fatalf("expected namesToIds to be %+v but got %+v", expectedNamesToIds, *actualNamesToIds)
	}

	if !reflect.DeepEqual(expectedUrisToNames, *actualUrisToNames) {
		t.Fatalf("expected urisToNames to be %+v but got %+v", expectedUrisToNames, *actualUrisToNames)
	}
}

func TestResourceIDNamingContainingAConstantAndSuffix(t *testing.T) {
	dnsResourceId := models.ParsedResourceId{
		Constants: map[string]models.ConstantDetails{
			"DnsRecordType": {
				FieldType: models.StringConstant,
				Values: map[string]string{
					"A":    "A",
					"AAAA": "AAAA",
				},
			},
		},
		Segments: []models.ResourceIdSegment{
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("subscriptions"),
				Name:       "subscriptions",
			},
			{
				Name: "subscriptionId",
				Type: models.SubscriptionIdSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("resourceGroups"),
				Name:       "resourceGroups",
			},
			{
				Name: "resourceGroupName",
				Type: models.ResourceGroupSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("providers"),
				Name:       "providers",
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("Microsoft.Network"),
				Name:       "Microsoft.Network",
			},
			{
				Type:              models.ConstantSegment,
				Name:              "recordType",
				ConstantReference: strPtr("DnsRecordType"),
			},
			{
				Name: "recordName",
				Type: models.UserSpecifiedSegment,
			},
		},
	}

	input := map[string]resourceUriMetadata{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/{recordType}/{recordName}/forceUpdate": {
			resourceIdName: nil,
			resourceId:     &dnsResourceId,
			uriSuffix:      strPtr("/forceUpdate"),
		},
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/{recordType}/{recordName}": {
			resourceIdName: nil,
			resourceId:     &dnsResourceId,
			uriSuffix:      nil,
		},
	}
	expectedNamesToIds := map[string]models.ParsedResourceId{
		"RecordTypeId": dnsResourceId,
	}
	expectedUrisToNames := map[string]string{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/{recordType}/{recordName}": "RecordTypeId",
	}

	actualNamesToIds, actualUrisToNames, err := determineNamesForResourceIds(input)
	if err != nil {
		t.Fatalf("error: %+v", err)
		return
	}

	if !reflect.DeepEqual(expectedNamesToIds, *actualNamesToIds) {
		t.Fatalf("expected namesToIds to be %+v but got %+v", expectedNamesToIds, *actualNamesToIds)
	}

	if !reflect.DeepEqual(expectedUrisToNames, *actualUrisToNames) {
		t.Fatalf("expected urisToNames to be %+v but got %+v", expectedUrisToNames, *actualUrisToNames)
	}
}

func TestResourceIdNamingTopLevelResourceId(t *testing.T) {
	input := map[string]resourceUriMetadata{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}": {
			resourceIdName: nil,
			resourceId:     &virtualMachineResourceId,
			uriSuffix:      nil,
		},
	}
	expectedNamesToIds := map[string]models.ParsedResourceId{
		"VirtualMachineId": virtualMachineResourceId,
	}
	expectedUrisToNames := map[string]string{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}": "VirtualMachineId",
	}

	actualNamesToIds, actualUrisToNames, err := determineNamesForResourceIds(input)
	if err != nil {
		t.Fatalf("error: %+v", err)
		return
	}

	if !reflect.DeepEqual(expectedNamesToIds, *actualNamesToIds) {
		t.Fatalf("expected namesToIds to be %+v but got %+v", expectedNamesToIds, *actualNamesToIds)
	}

	if !reflect.DeepEqual(expectedUrisToNames, *actualUrisToNames) {
		t.Fatalf("expected urisToNames to be %+v but got %+v", expectedUrisToNames, *actualUrisToNames)
	}
}

func TestResourceIdNamingTopLevelAndNestedResourceId(t *testing.T) {
	input := map[string]resourceUriMetadata{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}": {
			resourceIdName: nil,
			resourceId:     &virtualMachineResourceId,
			uriSuffix:      nil,
		},
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}/extensions/{extensionName}": {
			resourceIdName: nil,
			resourceId:     &virtualMachineExtensionResourceId,
			uriSuffix:      nil,
		},
	}
	expectedNamesToIds := map[string]models.ParsedResourceId{
		"VirtualMachineId": virtualMachineResourceId,
		"ExtensionId":      virtualMachineExtensionResourceId,
	}
	expectedUrisToNames := map[string]string{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}":                            "VirtualMachineId",
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}/extensions/{extensionName}": "ExtensionId",
	}

	actualNamesToIds, actualUrisToNames, err := determineNamesForResourceIds(input)
	if err != nil {
		t.Fatalf("error: %+v", err)
		return
	}

	if !reflect.DeepEqual(expectedNamesToIds, *actualNamesToIds) {
		t.Fatalf("expected namesToIds to be %+v but got %+v", expectedNamesToIds, *actualNamesToIds)
	}

	if !reflect.DeepEqual(expectedUrisToNames, *actualUrisToNames) {
		t.Fatalf("expected urisToNames to be %+v but got %+v", expectedUrisToNames, *actualUrisToNames)
	}
}

func TestResourceIdNamingNestedResourceId(t *testing.T) {
	input := map[string]resourceUriMetadata{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}/extensions/{extensionName}": {
			resourceIdName: nil,
			resourceId:     &virtualMachineExtensionResourceId,
			uriSuffix:      nil,
		},
	}
	expectedNamesToIds := map[string]models.ParsedResourceId{
		"ExtensionId": virtualMachineExtensionResourceId,
	}
	expectedUrisToNames := map[string]string{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}/extensions/{extensionName}": "ExtensionId",
	}

	actualNamesToIds, actualUrisToNames, err := determineNamesForResourceIds(input)
	if err != nil {
		t.Fatalf("error: %+v", err)
		return
	}

	if !reflect.DeepEqual(expectedNamesToIds, *actualNamesToIds) {
		t.Fatalf("expected namesToIds to be %+v but got %+v", expectedNamesToIds, *actualNamesToIds)
	}

	if !reflect.DeepEqual(expectedUrisToNames, *actualUrisToNames) {
		t.Fatalf("expected urisToNames to be %+v but got %+v", expectedUrisToNames, *actualUrisToNames)
	}
}

func TestResourceIdNamingResourceUnderScope(t *testing.T) {
	input := map[string]resourceUriMetadata{
		"/{scope}/providers/Microsoft.Monitor/extensions/{extensionName}": {
			resourceIdName: nil,
			resourceId:     &scopedMonitorResourceId,
			uriSuffix:      nil,
		},
	}
	expectedNamesToIds := map[string]models.ParsedResourceId{
		"ScopedExtensionId": scopedMonitorResourceId,
	}
	expectedUrisToNames := map[string]string{
		"/{scope}/providers/Microsoft.Monitor/extensions/{extensionName}": "ScopedExtensionId",
	}

	actualNamesToIds, actualUrisToNames, err := determineNamesForResourceIds(input)
	if err != nil {
		t.Fatalf("error: %+v", err)
		return
	}

	if !reflect.DeepEqual(expectedNamesToIds, *actualNamesToIds) {
		t.Fatalf("expected namesToIds to be %+v but got %+v", expectedNamesToIds, *actualNamesToIds)
	}

	if !reflect.DeepEqual(expectedUrisToNames, *actualUrisToNames) {
		t.Fatalf("expected urisToNames to be %+v but got %+v", expectedUrisToNames, *actualUrisToNames)
	}
}

func TestResourceIdNamingConflictingTwoLevels(t *testing.T) {
	input := map[string]resourceUriMetadata{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/extensions/{extensionName}": {
			resourceIdName: nil,
			resourceId:     &virtualNetworkExtensionResourceId,
			uriSuffix:      nil,
		},
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}/extensions/{extensionName}": {
			resourceIdName: nil,
			resourceId:     &virtualMachineExtensionResourceId,
			uriSuffix:      nil,
		},
	}
	expectedNamesToIds := map[string]models.ParsedResourceId{
		"VirtualMachineExtensionId": virtualMachineExtensionResourceId,
		"ExtensionId":               virtualNetworkExtensionResourceId,
	}
	expectedUrisToNames := map[string]string{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}/extensions/{extensionName}": "VirtualMachineExtensionId",
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/extensions/{extensionName}":                                      "ExtensionId",
	}

	actualNamesToIds, actualUrisToNames, err := determineNamesForResourceIds(input)
	if err != nil {
		t.Fatalf("error: %+v", err)
		return
	}

	if !reflect.DeepEqual(expectedNamesToIds, *actualNamesToIds) {
		t.Fatalf("expected namesToIds to be:\n\n%+v\nbut got:\n\n%+v", expectedNamesToIds, *actualNamesToIds)
	}

	if !reflect.DeepEqual(expectedUrisToNames, *actualUrisToNames) {
		t.Fatalf("expected urisToNames to be\n\n%+v\nbut got:\n\n%+v", expectedUrisToNames, *actualUrisToNames)
	}
}

func TestResourceIdNamingConflictingMultipleLevels(t *testing.T) {
	workerPoolInstanceResourceId := models.ParsedResourceId{
		Constants: map[string]models.ConstantDetails{},
		Segments: []models.ResourceIdSegment{
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("subscriptions"),
				Name:       "subscriptions",
			},
			{
				Name: "subscriptionId",
				Type: models.SubscriptionIdSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("resourceGroups"),
				Name:       "resourceGroups",
			},
			{
				Name: "resourceGroupName",
				Type: models.ResourceGroupSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("providers"),
				Name:       "providers",
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("Microsoft.Web"),
				Name:       "Microsoft.Web",
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("hostingEnvironments"),
				Name:       "hostingEnvironments",
			},
			{
				Name: "name",
				Type: models.UserSpecifiedSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("workerPools"),
				Name:       "workerPools",
			},
			{
				Name: "workerPoolName",
				Type: models.UserSpecifiedSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("instances"),
				Name:       "instances",
			},
			{
				Name: "instance",
				Type: models.UserSpecifiedSegment,
			},
		},
	}
	multiRolePoolInstanceResourceId := models.ParsedResourceId{
		Constants: map[string]models.ConstantDetails{},
		Segments: []models.ResourceIdSegment{
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("subscriptions"),
				Name:       "subscriptions",
			},
			{
				Name: "subscriptionId",
				Type: models.SubscriptionIdSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("resourceGroups"),
				Name:       "resourceGroups",
			},
			{
				Name: "resourceGroupName",
				Type: models.ResourceGroupSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("providers"),
				Name:       "providers",
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("Microsoft.Web"),
				Name:       "Microsoft.Web",
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("hostingEnvironments"),
				Name:       "hostingEnvironments",
			},
			{
				Name: "name",
				Type: models.UserSpecifiedSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("multiRolePools"),
				Name:       "multiRolePools",
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("default"),
				Name:       "default",
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("instances"),
				Name:       "instances",
			},
			{
				Name: "instance",
				Type: models.UserSpecifiedSegment,
			},
		},
	}
	slotInstanceProcessModuleResourceId := models.ParsedResourceId{
		Constants: map[string]models.ConstantDetails{},
		Segments: []models.ResourceIdSegment{
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("subscriptions"),
				Name:       "subscriptions",
			},
			{
				Name: "subscriptionId",
				Type: models.SubscriptionIdSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("resourceGroups"),
				Name:       "resourceGroups",
			},
			{
				Name: "resourceGroupName",
				Type: models.ResourceGroupSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("providers"),
				Name:       "providers",
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("Microsoft.Web"),
				Name:       "Microsoft.Web",
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("sites"),
				Name:       "sites",
			},
			{
				Name: "name",
				Type: models.UserSpecifiedSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("slots"),
				Name:       "slots",
			},
			{
				Name: "slot",
				Type: models.UserSpecifiedSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("instances"),
				Name:       "instances",
			},
			{
				Name: "instanceId",
				Type: models.UserSpecifiedSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("processes"),
				Name:       "processes",
			},
			{
				Name: "processId",
				Type: models.UserSpecifiedSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("modules"),
				Name:       "modules",
			},
			{
				Name: "baseAddress",
				Type: models.UserSpecifiedSegment,
			},
		},
	}
	instanceProcessModuleResourceId := models.ParsedResourceId{
		Constants: map[string]models.ConstantDetails{},
		Segments: []models.ResourceIdSegment{
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("subscriptions"),
				Name:       "subscriptions",
			},
			{
				Name: "subscriptionId",
				Type: models.SubscriptionIdSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("resourceGroups"),
				Name:       "resourceGroups",
			},
			{
				Name: "resourceGroupName",
				Type: models.ResourceGroupSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("providers"),
				Name:       "providers",
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("Microsoft.Web"),
				Name:       "Microsoft.Web",
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("sites"),
				Name:       "sites",
			},
			{
				Name: "name",
				Type: models.UserSpecifiedSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("instances"),
				Name:       "instances",
			},
			{
				Name: "instanceId",
				Type: models.UserSpecifiedSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("processes"),
				Name:       "processes",
			},
			{
				Name: "processId",
				Type: models.UserSpecifiedSegment,
			},
			{
				Type:       models.StaticSegment,
				FixedValue: strPtr("modules"),
				Name:       "modules",
			},
			{
				Name: "baseAddress",
				Type: models.UserSpecifiedSegment,
			},
		},
	}

	input := map[string]resourceUriMetadata{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/instances/{instance}": {
			resourceIdName: nil,
			resourceId:     &workerPoolInstanceResourceId,
			uriSuffix:      nil,
		},
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/instances/{instance}": {
			resourceIdName: nil,
			resourceId:     &multiRolePoolInstanceResourceId,
			uriSuffix:      nil,
		},
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/modules/{baseAddress}": {
			resourceIdName: nil,
			resourceId:     &slotInstanceProcessModuleResourceId,
			uriSuffix:      nil,
		},
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/modules/{baseAddress}": {
			resourceIdName: nil,
			resourceId:     &instanceProcessModuleResourceId,
			uriSuffix:      nil,
		},
	}
	expectedNamesToIds := map[string]models.ParsedResourceId{
		"WorkerPoolInstanceId": workerPoolInstanceResourceId,
		"InstanceId":           multiRolePoolInstanceResourceId,
		"ProcessModuleId":      slotInstanceProcessModuleResourceId,
		"ModuleId":             instanceProcessModuleResourceId,
	}
	expectedUrisToNames := map[string]string{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/instances/{instance}":                 "WorkerPoolInstanceId",
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/instances/{instance}":                       "InstanceId",
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/modules/{baseAddress}": "ProcessModuleId",
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/modules/{baseAddress}":              "ModuleId",
	}

	actualNamesToIds, actualUrisToNames, err := determineNamesForResourceIds(input)
	if err != nil {
		t.Fatalf("error: %+v", err)
		return
	}

	if !reflect.DeepEqual(expectedNamesToIds, *actualNamesToIds) {
		t.Fatalf("expected namesToIds to be:\n\n%+v\n\nbut got:\n%+v", expectedNamesToIds, *actualNamesToIds)
	}

	if !reflect.DeepEqual(expectedUrisToNames, *actualUrisToNames) {
		t.Fatalf("expected urisToNames to be:\n\n%+v\n\nbut got:\n%+v", expectedUrisToNames, *actualUrisToNames)
	}
}

func TestResourceIdNamingSignalRId(t *testing.T) {
	input := map[string]resourceUriMetadata{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/SignalR/{resourceName}": {
			resourceIdName: nil,
			resourceId:     &signalRResourceId,
			uriSuffix:      nil,
		},
	}
	expectedNamesToIds := map[string]models.ParsedResourceId{
		"SignalRId": signalRResourceId,
	}
	expectedUrisToNames := map[string]string{
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/SignalR/{resourceName}": "SignalRId",
	}

	actualNamesToIds, actualUrisToNames, err := determineNamesForResourceIds(input)
	if err != nil {
		t.Fatalf("error: %+v", err)
		return
	}

	if !reflect.DeepEqual(expectedNamesToIds, *actualNamesToIds) {
		t.Fatalf("expected namesToIds to be %+v but got %+v", expectedNamesToIds, *actualNamesToIds)
	}

	if !reflect.DeepEqual(expectedUrisToNames, *actualUrisToNames) {
		t.Fatalf("expected urisToNames to be %+v but got %+v", expectedUrisToNames, *actualUrisToNames)
	}
}

func strPtr(in string) *string {
	return &in
}