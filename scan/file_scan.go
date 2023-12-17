package scan

import (
	"line-by-line/param"
	"strings"
)

func ScanForErrors(lineValue string, lineIndex uint8, valuesToSearchFor ...string) map[string]param.Errors {
	var errors = make(map[string]param.Errors)
	for _, value := range valuesToSearchFor {
		if strings.Contains(lineValue, value) {
			if foundError, ok := errors[value]; ok {
				errors[value] = *foundError.UpdateCount().UpdateLines(lineIndex)
			} else {
				errors[value] = param.New(lineIndex)
			}
		}
	}
	return errors
}
