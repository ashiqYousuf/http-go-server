package protocol

type Method int

const (
	MethodGet    Method = 0
	MethodPost   Method = 1
	MethodDelete Method = 2
	MethodPut    Method = 3
	MethodPatch  Method = 4
)

func (m Method) String() string {
	switch m {
	case MethodGet:
		return "GET"
	case MethodPost:
		return "POST"
	case MethodDelete:
		return "DELETE"
	case MethodPut:
		return "PUT"
	case MethodPatch:
		return "PUT"
	default:
		return "Unknown Method"
	}
}
