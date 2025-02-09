package http_utils

import (
	"fmt"
	"strings"

	"github.com/ashiqYousuf/http-go-server/internal/constants"
	"github.com/ashiqYousuf/http-go-server/internal/status"
	"github.com/ashiqYousuf/http-go-server/utils"
)

func GetRequestLine(request string) string {
	requestSlice := strings.Split(request, "\r\n\r\n")
	return strings.Split(requestSlice[0], "\r\n")[0]
}

func GetRequestHeaders(request string) map[string]string {
	requestSlice := strings.Split(request, "\r\n\r\n")
	headerSlice := strings.Split(requestSlice[0], "\r\n")[1:]
	return utils.SliceToMap(headerSlice, ":")
}

func GetRequestBody(request string) string {
	return strings.Split(request, "\r\n\r\n")[1]
}

func GetLineRequest(status status.Status) string {
	return fmt.Sprintf("%s %d %s%s", constants.HTTP_VERSION, status, status.String(), constants.CRLF)
}
