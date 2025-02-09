package status

// Status type represents HTTP status codes.
type Status int

const (
	StatusOK        Status = 200
	StatusCreated   Status = 201
	StatusAccepted  Status = 202
	StatusNoContent Status = 204

	StatusBadRequest       Status = 400
	StatusUnauthorized     Status = 401
	StatusForbidden        Status = 403
	StatusNotFound         Status = 404
	StatusMethodNotAllowed Status = 405

	StatusInternalServerError Status = 500
	StatusNotImplemented      Status = 501
	StatusBadGateway          Status = 502
	StatusServiceUnavailable  Status = 503
	StatusGatewayTimeout      Status = 504
)

func (s Status) String() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusCreated:
		return "Created"
	case StatusAccepted:
		return "Accepted"
	case StatusNoContent:
		return "No Content"

	case StatusBadRequest:
		return "Bad Request"
	case StatusUnauthorized:
		return "Unauthorized"
	case StatusForbidden:
		return "Forbidden"
	case StatusNotFound:
		return "Not Found"
	case StatusMethodNotAllowed:
		return "Method Not Allowed"

	case StatusInternalServerError:
		return "Internal Server Error"
	case StatusNotImplemented:
		return "Not Implemented"
	case StatusBadGateway:
		return "Bad Gateway"
	case StatusServiceUnavailable:
		return "Service Unavailable"
	case StatusGatewayTimeout:
		return "Gateway Timeout"
	default:
		return "Unknown Status"
	}
}
