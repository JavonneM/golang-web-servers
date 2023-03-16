package apperrors

import (
    "runtime"
)

const (
    GatewayErrorType = iota
    ServiceErrorType
    RouteErrorType
)

type BaseErrorInterface interface {
    Error() string
    ErrorType() int 
    ErrorMessage() string
    ErrorStatusCode() int
}

type BaseErrorStruct struct {
    rawError error
    errorType int
    message string
    httpCode int
}

func (error BaseErrorStruct) Error() string {
    return error.rawError.Error()
}
func (error BaseErrorStruct) ErrorType() int {
    return error.errorType
}
func (error BaseErrorStruct) ErrorMessage() string {
    return error.message
}
func (error BaseErrorStruct) ErrorStatusCode() int {
    return error.httpCode
}


func dumpStack() string {
    buf := make([]byte, 1<<16)
    runtime.Stack(buf, true)
    return string(buf)
}


