package errors

import "net/http"

type ApiErrorInterface interface {
	Status() int
	Message() string
	Error() string
}

type ApiErrorStruct struct {
	status  int    `json:"status"`
	message string `json:"message"`
	error   string `json:"errors,omitempty"`
}

func (e *ApiErrorStruct) Status() int {
	return e.status
}

func (e *ApiErrorStruct) Message() string {
	return e.message
}

func (e *ApiErrorStruct) Error() string {
	return e.error
}

func NewApiError(statusCode int, message string) ApiErrorInterface {
	return &ApiErrorStruct{
		status:  statusCode,
		message: message,
	}
}

func NewInternalServerApiError(message string) ApiErrorInterface {
	return &ApiErrorStruct{
		status:  http.StatusInternalServerError,
		message: message,
	}
}

func NewNotFoundApiError(message string) ApiErrorInterface {
	return &ApiErrorStruct{
		status:  http.StatusNotFound,
		message: message,
	}
}

func NewBadRequestApiError(message string) ApiErrorInterface {
	return &ApiErrorStruct{
		status:  http.StatusBadRequest,
		message: message,
	}
}
