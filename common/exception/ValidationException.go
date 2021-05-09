package exception

type ValidationException struct {
	error string
	code  int
}

func (a *ValidationException) Code() int {
	return a.code
}

func (a *ValidationException) SetCode(code int) {
	a.code = code
}

func (a *ValidationException) Error() string {
	return a.error
}
func CreateValidationException(error string) *ValidationException {
	return &ValidationException{
		code:  400,
		error: error,
	}
}
