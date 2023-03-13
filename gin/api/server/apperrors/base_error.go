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
    RawErrorObj() error 
    ErrorType() int 
    ErrorMesage() string
}

type BaseErrorStruct struct {

    rawError error
    errorType int
    message string
}
type GatewayErrorStruct struct {
    BaseErrorStruct
    domain string
}
func dumpStack() string {
    buf := make([]byte, 1<<16)
    runtime.Stack(buf, true)
    return string(buf)
}

func NewGatewayError(errorText string, rawError error, domain string) BaseError {
    return GatewayErrorStruct {
        rawError: rawError, 
        errorType: GatewayErrorType,
        message: errorText,
        domain: domain,
    }
}


func NewServiceError(errorText string, rawError error, domain string) BaseError {
    return GatewayErrorStruct {
        rawError: rawError, 
        errorType: GatewayErrorType,
        message: errorText,
        domain: domain,
    }
}
