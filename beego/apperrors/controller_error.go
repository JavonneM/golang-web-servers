package apperrors

import (
)
const (
   DataInconsistencyError = iota

)

type ControllerErrorStruct struct {
    BaseErrorStruct
    domain string
}

type ControllerErrorInterface interface {
    BaseErrorInterface
    FindDomainErrors() string
}

func (error ControllerErrorStruct) FindDomainErrors() string {
    return "FindDomainErrors method"
}


func NewControllerError(errorText string, rawError error, domain string, httpStatusCode int) ControllerErrorInterface {
    return ControllerErrorStruct {
        BaseErrorStruct{
            rawError: rawError, 
            errorType: ServiceErrorType,
            message: errorText,
            httpCode: httpStatusCode,
        },
        domain,
    }
}
