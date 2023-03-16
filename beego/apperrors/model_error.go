package apperrors

type ModelErrorInterface interface {
    BaseErrorInterface
    GetDomain() string
}
type ModelErrorStruct struct {
    BaseErrorStruct
    domain string
}
func (error ModelErrorStruct) GetDomain() string {
    return error.domain
}

func NewModelError(statusCode int, errorText string, rawError error, domain string) ModelErrorInterface {
    return ModelErrorStruct {
        BaseErrorStruct: BaseErrorStruct{
            rawError: rawError, 
            errorType: GatewayErrorType,
            message: errorText,
            httpCode: statusCode,
        },
        domain: domain,
    }
}

