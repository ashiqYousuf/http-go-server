package utils

import (
	"strings"
)

func SliceToMap(slice []string, sep string) map[string]string {
	resultMap := make(map[string]string, len(slice))

	for _, entry := range slice {
		entryParts := strings.Split(entry, sep)
		if len(entryParts) > 2 {
			resultMap[entryParts[0]] = entryParts[1]
		}
	}
	return resultMap
}
