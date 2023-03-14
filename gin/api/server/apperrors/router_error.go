package apperrors

type RouterErrorStruct struct {
    BaseErrorStruct
}


func NewRouterError(errorText string, rawError error, domain string, httpStatusCode int) ServiceErrorStruct {
    return ServiceErrorStruct {
        BaseErrorStruct{
            rawError: rawError, 
            errorType: RouteErrorType,
            message: errorText,
            httpCode: httpStatusCode,
        },
        domain,
    }
}
