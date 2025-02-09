package server

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	DIR = "./data"
)

func (server *Server) routeRequests(httpRequest *HTTPRequest) *HTTPResponse {
	httpResponse := &HTTPResponse{}

	path := httpRequest.URL.Path
	if path == "/" || path == "" {
		httpResponse = NewHTTPResponse("", "", nil)
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
	} else if strings.HasPrefix(path, "/files/") {
		fileName := strings.TrimPrefix(path, "/files/")
		filePath := fmt.Sprintf("%s/%s", DIR, fileName)
		fmt.Println(filePath, fileName)
		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("error while reading file", err)
			httpResponse = NewHTTPResponse("HTTP/1.1 404 Not Found", "", nil)
		} else {
			httpResponse = NewHTTPResponse("HTTP/1.1 200 OK", string(data), map[string]string{
				"Content-Type":   "application/octet-stream",
				"Content-Length": strconv.Itoa(len(data)),
			})
		}
	} else if httpRequest.URL.Method == "POST" && strings.HasPrefix(path, "/echo") {
		httpResponse = NewHTTPResponse("HTTP/1.1 200 OK", httpRequest.Body, map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": strconv.Itoa(len(httpRequest.Body)),
		})
	} else {
		httpResponse = NewHTTPResponse("HTTP/1.1 404 Not Found", "", nil)
	}

	return httpResponse
}
