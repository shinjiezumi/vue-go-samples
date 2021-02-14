package common

type ApplicationError struct {
	HttpCode   int
	Message    Message
	LogMessage string
}

func NewApplicationError(code int, msg Message) *ApplicationError {
	return &ApplicationError{
		HttpCode: code,
		Message:  msg,
	}
}

func (e ApplicationError) Error() string {
	return string(e.Message)
}
