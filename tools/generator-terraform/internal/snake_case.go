package internal

import "strings"

func SnakeCase(input string) string {
	output := ""

	for i, c := range input {
		rawVal := string(c)
		if rawVal == strings.ToUpper(rawVal) {
			if i != 0 {
				output += "_"
			}

			rawVal = strings.ToLower(rawVal)
		}

		output += rawVal
	}

	return output
}
