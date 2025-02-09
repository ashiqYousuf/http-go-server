package protocol

import "github.com/ashiqYousuf/http-go-server/internal/status"

type HTTPResponse struct {
	*HTTPProtocol
	StatusCode status.Status
}

func NewHTTPResponse(httpProtocol *HTTPProtocol, status status.Status) *HTTPResponse {
	return &HTTPResponse{
		HTTPProtocol: httpProtocol,
		StatusCode:   status,
	}
}
