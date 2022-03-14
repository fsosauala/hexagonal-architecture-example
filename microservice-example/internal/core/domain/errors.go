package domain

import (
	"fmt"
	"net/http"
)

var AlreadyExistsError = fmt.Errorf("the country already exists")

type CustomErr struct {
	HTTPCode  int    `json:"-"`
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

var (
	ErrUnknownError = CustomErr{
		HTTPCode:  http.StatusInternalServerError,
		Message:   "Something unexpected happened",
		ErrorCode: 1,
	}
	ErrCannotParseBody = CustomErr{
		HTTPCode:  http.StatusBadRequest,
		Message:   "Error parsing body",
		ErrorCode: 2,
	}
	ErrEmptyName = CustomErr{
		HTTPCode:  http.StatusBadRequest,
		Message:   "The country name cannot be empty",
		ErrorCode: 3,
	}
	ErrBadRequest = CustomErr{
		HTTPCode:  http.StatusBadRequest,
		Message:   "The country already exists",
		ErrorCode: 4,
	}
)

func (ce CustomErr) Error() string {
	return ce.Message
}
