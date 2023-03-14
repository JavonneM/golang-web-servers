package apperrors

import (
)
const (
   DataInconsistencyError = iota

)

type ServiceErrorStruct struct {
    BaseErrorStruct
    domain string
}

type ServiceErrorInterface interface {
    BaseErrorInterface
    FindDomainErrors() string
}

func (error ServiceErrorStruct) FindDomainErrors() string {
    return "FindDomainErrors method"
}


func NewServiceError(errorText string, rawError error, domain string, httpStatusCode int) ServiceErrorInterface {
    return ServiceErrorStruct {
        BaseErrorStruct{
            rawError: rawError, 
            errorType: ServiceErrorType,
            message: errorText,
            httpCode: httpStatusCode,
        },
        domain,
    }
}
