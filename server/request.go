package server

import "strings"

func (server *Server) parseRequest(buffer []byte, limit int) *HTTPProtocol {
	request := string(buffer[:limit])
	requestSlice := strings.Split(request, "\r\n\r\n")
	/*
		<Request Line>      --> e.g., "GET /path HTTP/1.1" \r\n
		<Headers>           --> e.g., "Host: 127.0.0.1" (seperated by \r\n)
		(blank line)        --> "\r\n" (End of headers)
		<Body (optional)>   --> (Present in POST, PUT, etc.)
	*/

	statusLine := strings.Split(requestSlice[0], "\r\n")[0]
	// Headers are seperate by \r\n
	headerSlice := strings.Split(requestSlice[0], "\r\n")[1:]
	requestBody := requestSlice[1]

	headerMap := server.getHeaderMap(headerSlice)
	return NewHTTPProtocol(statusLine, requestBody, headerMap)
}
