package httperrors

import (
	"fmt"

	"github.com/ashiqYousuf/http-go-server/internal/constants"
)

var (
	ErrBufferLimitExceeded = fmt.Errorf("buffer limit exceeded %d", constants.MAX_BUFFER_SIZE)
)
