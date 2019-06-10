package Errors

type OutOfRangeException struct {
	message string
	code    int
}

func NewOutOfRangeException(message string, code ...int) Exception {
	code = append(code, 0)
	return OutOfRangeException{
		message: message,
		code:    code[0],
	}
}

func (this OutOfRangeException) GetMessage() string {
	return this.message
}

func (this OutOfRangeException) Error() string {
	return this.GetMessage()
}

func (this OutOfRangeException) GetCode() int {
	return this.code
}
