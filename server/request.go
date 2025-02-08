package server

import (
	"strings"

	"github.com/ashiqYousuf/http-go-server/utils"
)

func (server *Server) parseRequest(buffer []byte, limit int) *HTTPRequest {
	request := string(buffer[:limit])
	requestSlice := strings.Split(request, "\r\n\r\n")
	statusLine := strings.Split(requestSlice[0], "\r\n")[0]
	headerSlice := strings.Split(requestSlice[0], "\r\n")[1:]
	requestBody := requestSlice[1]

	headerMap := utils.SliceToMap(headerSlice, ":")
	url := server.parseURL(statusLine)
	return NewHTTPRequest(statusLine, requestBody, headerMap, url)
}

func (server *Server) parseURL(statusLine string) *URL {
	statusSlice := strings.Split(statusLine, " ")

	requestMethod := statusSlice[0]
	requestProtocol := statusSlice[2]
	requestRoute := strings.Split(statusSlice[1], "?")

	path, queryParams := "", ""
	if len(requestRoute) > 1 {
		path, queryParams = requestRoute[0], requestRoute[1]
	} else {
		path = requestRoute[0]
	}

	queryParamSlice := strings.Split(queryParams, "&")
	queryParamMap := utils.SliceToMap(queryParamSlice, "=")

	return &URL{
		Method:      requestMethod,
		Protocol:    requestProtocol,
		Path:        path,
		QueryParams: queryParamMap,
	}
}
