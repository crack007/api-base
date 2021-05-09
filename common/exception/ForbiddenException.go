package exception

type ForbiddenException struct {
	code  int
	error string
}

func (f *ForbiddenException) Error() string {
	return f.error
}

func (f *ForbiddenException) SetError(error string) {
	f.error = error
}

func (f *ForbiddenException) Code() int {
	return f.code
}

func (f *ForbiddenException) SetCode(code int) {
	f.code = code
}

func CreateForbiddenException(error string) *ForbiddenException {
	return &ForbiddenException{
		code:  403,
		error: error,
	}
}
