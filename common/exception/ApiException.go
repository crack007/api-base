package exception

type ApiException struct {
	error string
	code  int
}

func (a *ApiException) Code() int {
	return a.code
}

func (a *ApiException) SetCode(code int) {
	a.code = code
}

func (a *ApiException) Error() string {
	return a.error
}
func CreateApiException(code int, error string) *ApiException {
	return &ApiException{
		code:  code,
		error: error,
	}
}
