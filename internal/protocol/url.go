package protocol

type URL struct {
	Path     string
	QueryMap map[string]string
	Method   string
	Protocol string
}

func NewURL(path string, method string, protocol string, queryParams map[string]string) *URL {
	return &URL{
		Path:     path,
		QueryMap: queryParams,
		Method:   method,
		Protocol: protocol,
	}
}
