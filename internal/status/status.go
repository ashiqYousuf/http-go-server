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
	default:
		return "Unknown Status"
	}
}
