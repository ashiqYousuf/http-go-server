package protocol

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/ashiqYousuf/http-go-server/internal/constants"
	"github.com/ashiqYousuf/http-go-server/pkg/utils"
)

func HTTPResponseToBytes(httpResponse *HTTPResponse) ([]byte, error) {
	buffer := new(bytes.Buffer)
	_, err := fmt.Fprintf(buffer, "%s", utils.GetLineRequest(httpResponse.StatusCode))
	if err != nil {
		return nil, err
	}

	for key, value := range httpResponse.Headers {
		_, err = fmt.Fprintf(buffer, "%s:%s%s", strings.ToTitle(key), value, constants.CRLF)
		if err != nil {
			return nil, err
		}
	}

	_, err = fmt.Fprintf(buffer, constants.CRLF) // mark the end of header section
	if err != nil {
		return nil, err
	}

	_, err = fmt.Fprintf(buffer, "%s", httpResponse.Body)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
