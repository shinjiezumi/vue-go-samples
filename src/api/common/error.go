package common

import "fmt"

type ApplicationError struct {
	HttpCode int
	Message  Message
}

func NewApplicationError(code int, msg Message, origErr error) *ApplicationError {
	if origErr != nil {
		fmt.Printf("[ERROR]%s\n", origErr.Error())
	}
	return &ApplicationError{
		HttpCode: code,
		Message:  msg,
	}
}

func (e ApplicationError) Error() string {
	return string(e.Message)
}
