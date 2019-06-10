package Errors

type OutOfBoundsException struct {
	message string
	code    int
}

func NewOutOfBoundsException(message string, code ...int) Exception {
	code = append(code, 0)
	return OutOfBoundsException{
		message: message,
		code:    code[0],
	}
}

func (this OutOfBoundsException) GetMessage() string {
	return this.message
}

func (this OutOfBoundsException) Error() string {
	return this.GetMessage()
}

func (this OutOfBoundsException) GetCode() int {
	return this.code
}
