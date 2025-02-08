package server

type HTTPProtocol struct {
	Status string
	Body   string
	Header map[string]string
}

type HTTPResponse struct {
	HTTPProtocol
}

type HTTPRequest struct {
	HTTPProtocol
	URL *URL
}

type URL struct {
	Method      string
	Path        string
	QueryParams map[string]string
	Protocol    string
}

func NewHTTPProtocol(status, body string, header map[string]string) *HTTPProtocol {
	return &HTTPProtocol{
		Status: status,
		Body:   body,
		Header: header,
	}
}

func NewHTTPRequest(status, body string, header map[string]string, url *URL) *HTTPRequest {
	return &HTTPRequest{
		HTTPProtocol: *NewHTTPProtocol(status, body, header),
		URL:          url,
	}
}

func NewHTTPResponse(status, body string, header map[string]string) *HTTPResponse {
	return &HTTPResponse{
		HTTPProtocol: *NewHTTPProtocol(status, body, header),
	}
}
