package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ApiErrorInterface interface {
	Status() int
	Message() string
	Error() string
}

type ApiErrorStruct struct {
	Estatus  int    `json:"status"`
	Emessage string `json:"message"`
	Eerror   string `json:"errors,omitempty"`
}

func (e *ApiErrorStruct) Status() int {
	return e.Estatus
}

func (e *ApiErrorStruct) Message() string {
	return e.Emessage
}

func (e *ApiErrorStruct) Error() string {
	return e.Eerror
}

func NewApiError(statusCode int, message string) ApiErrorInterface {
	return &ApiErrorStruct{
		Estatus:  statusCode,
		Emessage: message,
	}
}

func NewApiErrorFromBytes(body []byte) (ApiErrorInterface, error) {
	var result ApiErrorStruct
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("invalid json for creating")
	}
	return &result, nil
}

func NewInternalServerApiError(message string) ApiErrorInterface {
	return &ApiErrorStruct{
		Estatus:  http.StatusInternalServerError,
		Emessage: message,
	}
}

func NewNotFoundApiError(message string) ApiErrorInterface {
	return &ApiErrorStruct{
		Estatus:  http.StatusNotFound,
		Emessage: message,
	}
}

func NewBadRequestApiError(message string) ApiErrorInterface {
	return &ApiErrorStruct{
		Estatus:  http.StatusBadRequest,
		Emessage: message,
	}
}
