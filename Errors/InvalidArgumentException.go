package Errors

type InvalidArgumentException struct {
	message string
	code    int
}

func NewInvalidArgumentException(message string, code ...int) Exception {
	code = append(code, 0)
	return InvalidArgumentException{
		message: message,
		code:    code[0],
	}
}

func (this InvalidArgumentException) GetMessage() string {
	return this.message
}

func (this InvalidArgumentException) Error() string {
	return this.GetMessage()
}

func (this InvalidArgumentException) GetCode() int {
	return this.code
}
