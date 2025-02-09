package service

import "github.com/ashiqYousuf/http-go-server/server"

func Home(httpRequest *server.HTTPRequest) *server.HTTPResponse {
	if httpRequest.URL.Method != "GET" {
		return server.NewHTTPResponse("HTTP/1.1 405 Method Not Allowed", "", nil)
	}
	return server.NewHTTPResponse("HTTP/1.1 200 OK", "", nil)
}

func File(httpRequest *server.HTTPRequest) *server.HTTPResponse {
	return nil
}

// func EchoGet(httpRequest *server.HTTPRequest) *server.HTTPResponse {
// 	if httpRequest.URL.Method != "GET" {
// return server.NewHTTPResponse("HTTP/1.1 405 Method Not Allowed", "", nil)
// 	}

// 	return server.NewHTTPResponse("HTTP/1.1 200 OK", "", map[string]string{
// 		"Content-Type": "text/plain",
// 	})
// }
