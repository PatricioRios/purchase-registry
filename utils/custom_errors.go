package utils

import "net/http"

// custom error, for service
type SrvcError interface {
	Error() string
	Code() int
}

type SrvcErrorImpl struct {
	CodeError int
	Message   string
}

func (e SrvcErrorImpl) Error() string {
	return e.Message
}

func (e SrvcErrorImpl) Code() int {
	return e.CodeError
}

func NewBadRequest(message string) SrvcError {
	return SrvcErrorImpl{
		CodeError: http.StatusBadRequest,
		Message:   message,
	}
}
func NewError(errorCode int, message string) SrvcError {
	return SrvcErrorImpl{
		CodeError: errorCode,
		Message:   message,
	}
}
