package server

import (
	"strconv"
	"strings"
)

func (server *Server) routeRequests(httpRequest *HTTPRequest) *HTTPResponse {
	httpResponse := &HTTPResponse{}

	path := httpRequest.URL.Path
	if path == "/" || path == "" {
		httpResponse = NewHTTPResponse("HTTP/1.1 200 OK", "", nil)
	} else if strings.HasPrefix(path, "/echo/") {
		data := path[6:]
		httpResponse = NewHTTPResponse("HTTP/1.1 200 OK", data, map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": strconv.Itoa(len(data)),
		})
	} else if path == "/user-agent" {
		userAgent := httpRequest.Header["User-Agent"]
		httpResponse = NewHTTPResponse("HTTP/1.1 200 OK", userAgent, map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": strconv.Itoa(len(userAgent)),
		})
	} else {
		httpResponse = NewHTTPResponse("HTTP/1.1 404 Not Found", "", nil)
	}

	return httpResponse
}
