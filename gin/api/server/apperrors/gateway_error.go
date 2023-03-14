package apperrors

type GatewayErrorInterface interface {
    BaseErrorInterface
    GetDomain() string
}
type GatewayErrorStruct struct {
    BaseErrorStruct
    domain string
}
func (error GatewayErrorStruct) GetDomain() string {
    return error.domain
}

func NewGatewayError(statusCode int, errorText string, rawError error, domain string) GatewayErrorInterface {
    return GatewayErrorStruct {
        BaseErrorStruct: BaseErrorStruct{
            rawError: rawError, 
            errorType: GatewayErrorType,
            message: errorText,
            httpCode: statusCode,
        },
        domain: domain,
    }
}

