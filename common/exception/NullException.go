package exception

type NullException struct {
	code  int
	error string
}

func (n *NullException) Error() string {
	return n.error
}

func (n *NullException) SetError(error string) {
	n.error = error
}

func (n *NullException) Code() int {
	return n.code
}

func (n *NullException) SetCode(code int) {
	n.code = code
}

func CreateNullException(error string) *NullException {
	return &NullException{
		code:  404,
		error: error,
	}
}
