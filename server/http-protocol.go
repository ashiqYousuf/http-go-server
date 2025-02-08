package server

type HTTPProtocol struct {
	Status string
	Body   string
	Header map[string]string
	URL    *URL
}

type URL struct {
	Method      string
	Path        string
	QueryParams map[string]string
	Protocol    string
}

func NewHTTPProtocol(status, body string, header map[string]string, url *URL) *HTTPProtocol {
	return &HTTPProtocol{
		Status: status,
		Body:   body,
		Header: header,
		URL:    url,
	}
}
