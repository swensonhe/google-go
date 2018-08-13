package google

type Error string

const (
	ErrInternal           = Error("internal_error")
	ErrInvalidCredentials = Error("invalid_credentials")
)

func (e Error) Error() string {
	return string(e)
}
