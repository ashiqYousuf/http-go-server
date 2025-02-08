package server

type HTTPProtocol struct {
	Status string
	Body   string
	Header map[string]string
	URL    URL
}

type URL struct {
	Path string
}

func NewHTTPProtocol(status, body string, header map[string]string) *HTTPProtocol {
	return &HTTPProtocol{
		Status: status,
		Body:   body,
		Header: header,
	}
}
