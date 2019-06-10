package Errors

type UnexpectedValueException struct {
	message string
	code    int
}

func NewUnexpectedValueException(message string, code ...int) Exception {
	code = append(code, 0)
	return UnexpectedValueException{
		message: message,
		code:    code[0],
	}
}

func (this UnexpectedValueException) GetMessage() string {
	return this.message
}

func (this UnexpectedValueException) Error() string {
	return this.GetMessage()
}

func (this UnexpectedValueException) GetCode() int {
	return this.code
}
