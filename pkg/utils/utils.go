package utils

import (
	"fmt"
	"io"
	"strings"

	"github.com/ashiqYousuf/http-go-server/internal/constants"
	"github.com/ashiqYousuf/http-go-server/internal/status"
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

func GetLineRequest(status status.Status) string {
	return fmt.Sprintf("%s %d %s%s", constants.HTTP_VERSION, status, status.String(), constants.CRLF)
}

func WriteBytes(writer io.Writer, body []byte) error {
	_, err := writer.Write(body)
	return err
}
