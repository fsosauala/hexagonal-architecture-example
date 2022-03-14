package domain

import (
	"encoding/json"
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
	ErrEmptyName = CustomErr{
		HTTPCode:  http.StatusBadRequest,
		Message:   "Te country name cannot be empty",
		ErrorCode: 1,
	}
	ErrBadRequest = CustomErr{
		HTTPCode:  http.StatusBadRequest,
		Message:   "The country already exists",
		ErrorCode: 2,
	}
	ErrUnknownError = CustomErr{
		HTTPCode:  http.StatusInternalServerError,
		Message:   "Something unexpected happened",
		ErrorCode: 3,
	}
)

func (ce CustomErr) Error() string {
	return ce.Message
}

func (ce CustomErr) String() string {
	data, err := json.Marshal(ce)
	if err != nil {
		return `{"message":"error unknown","errorCode":0}`
	}
	return string(data)
}
