package protocol

type HTTPRequest struct {
	*HTTPProtocol
	*URL
}

func NewHTTPRequest(httpProtocol *HTTPProtocol, url *URL) *HTTPRequest {
	return &HTTPRequest{
		HTTPProtocol: httpProtocol,
		URL:          url,
	}
}
