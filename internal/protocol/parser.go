package protocol

import (
	"net"
	"strings"

	"github.com/ashiqYousuf/http-go-server/internal/constants"
	"github.com/ashiqYousuf/http-go-server/internal/httperrors"
	"github.com/ashiqYousuf/http-go-server/pkg/validator"
	"github.com/ashiqYousuf/http-go-server/utils"
)

func ParseRequest(conn net.Conn) (*HTTPRequest, error) {
	buffer := make([]byte, constants.MAX_BUFFER_SIZE)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}

	if validator.BufferExceeded(n, constants.MAX_BUFFER_SIZE) {
		return nil, httperrors.ErrBufferLimitExceeded
	}

	request := string(buffer[:n])
	return parseRequest(request), nil
}

func parseRequest(request string) *HTTPRequest {
	requestLine := getRequestLine(request)
	headers := getRequestHeaders(request)
	body := getRequestBody(request)

	url := parseURL(requestLine)

	httpProtocol := NewHTTPProtocol(headers, body)
	return NewHTTPRequest(httpProtocol, url)
}

func getRequestLine(request string) string {
	requestSlice := strings.Split(request, "\r\n\r\n")
	return strings.Split(requestSlice[0], "\r\n")[0]
}

func getRequestHeaders(request string) map[string]string {
	requestSlice := strings.Split(request, "\r\n\r\n")
	headerSlice := strings.Split(requestSlice[0], "\r\n")[1:]
	return utils.SliceToMap(headerSlice, ":")
}

func getRequestBody(request string) string {
	return strings.Split(request, "\r\n\r\n")[1]
}

func parseURL(requestLine string) *URL {
	/*
		GET /path?q=123 HTTP/1.1
	*/
	statusSlice := strings.Split(requestLine, " ")

	requestMethod := statusSlice[0]
	requestProtocol := statusSlice[2]
	requestPath := strings.Split(statusSlice[1], "?")

	qParams := ""
	path := requestPath[0]

	if len(requestPath) == 2 {
		qParams = requestPath[1]
	}

	paramMap := utils.SliceToMap(strings.Split(qParams, "&"), "=")
	return &URL{
		Path:     path,
		Method:   requestMethod,
		Protocol: requestProtocol,
		QueryMap: paramMap,
	}
}
