package protocol

import (
	"net"
	"strings"

	"github.com/ashiqYousuf/http-go-server/internal/constants"
	"github.com/ashiqYousuf/http-go-server/internal/httperrors"
	http_utils "github.com/ashiqYousuf/http-go-server/internal/utils"
	"github.com/ashiqYousuf/http-go-server/pkg/utils"
	"github.com/ashiqYousuf/http-go-server/pkg/validator"
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
	requestLine := http_utils.GetRequestLine(request)
	headers := http_utils.GetRequestHeaders(request)
	body := http_utils.GetRequestBody(request)

	url := parseURL(requestLine)

	httpProtocol := NewHTTPProtocol(headers, body)
	return NewHTTPRequest(httpProtocol, url)
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
