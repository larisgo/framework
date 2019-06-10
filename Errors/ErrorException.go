package Errors

type ErrorException struct {
	message string
	code    int
}

func NewErrorException(message string, code ...int) Exception {
	code = append(code, 0)
	return ErrorException{
		message: message,
		code:    code[0],
	}
}

func (this ErrorException) GetMessage() string {
	return this.message
}

func (this ErrorException) Error() string {
	return this.GetMessage()
}

func (this ErrorException) GetCode() int {
	return this.code
}
