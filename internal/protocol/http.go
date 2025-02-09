package protocol

type HTTPProtocol struct {
	Headers map[string]string
	Body    string
}

func NewHTTPProtocol(headers map[string]string, body string) *HTTPProtocol {
	return &HTTPProtocol{
		Headers: headers,
		Body:    body,
	}
}
