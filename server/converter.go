package server

import (
	"bytes"
	"fmt"
	"strings"
)

func (server *Server) convertHTTPProtocolToBytes(httpProtocol *HTTPProtocol) ([]byte, error) {
	buffer := new(bytes.Buffer)

	_, err := fmt.Fprintf(buffer, "%s\r\n", httpProtocol.Status)
	if err != nil {
		return nil, err
	}

	for key, value := range httpProtocol.Header {
		_, err = fmt.Fprintf(buffer, "%s:%s\r\n", strings.ToTitle(key), value)
		if err != nil {
			return nil, err
		}
	}

	_, err = fmt.Fprintf(buffer, "\r\n") // mark the end of header section
	if err != nil {
		return nil, err
	}

	_, err = fmt.Fprintf(buffer, "%s", httpProtocol.Body)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
