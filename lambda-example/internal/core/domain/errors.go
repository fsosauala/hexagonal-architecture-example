package domain

import (
	"fmt"
	"net/http"
)

var AlreadyExistsError = fmt.Errorf("the country already exists")

type CustomErr struct {
	HTTPCode  int
	Message   string
	ErrorCode int
}

var (
	ErrEmptyName = CustomErr{
		HTTPCode:  http.StatusBadRequest,
		Message:   "Te country name cannot be empty",
		ErrorCode: 0,
	}
	ErrBadRequest = CustomErr{
		HTTPCode:  http.StatusBadRequest,
		Message:   "The country already exists",
		ErrorCode: 1,
	}
	ErrUnknownError = CustomErr{
		HTTPCode:  http.StatusInternalServerError,
		Message:   "Something unexpected happened",
		ErrorCode: 2,
	}
)

func (ce CustomErr) Error() string {
	return ce.Message
}
