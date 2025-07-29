package global

var (
	ErrBadRequest   = NewError(400, "bad request", nil)
	ErrUnauthorized = NewError(401, "unauthorized", nil)
	ErrForbidden    = NewError(403, "forbidden", nil)
)

type Error struct {
	Code     int
	Message  string
	Internal error
}

func (e Error) Error() string {
	return e.Message
}
func (e Error) Unwrap() error {
	return e.Internal
}

func NewError(code int, message string, internal error) Error {
	return Error{
		Code:     code,
		Message:  message,
		Internal: internal,
	}
}
