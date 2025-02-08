package utils

import "strings"

func SliceToMap(slice []string, seperator string) map[string]string {
	resultMap := make(map[string]string)
	for _, entry := range slice {
		entryParts := strings.Split(entry, seperator)
		if len(entryParts) >= 2 {
			resultMap[entryParts[0]] = strings.Join(entryParts[1:], "")
		}
	}

	return resultMap
}
