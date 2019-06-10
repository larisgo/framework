package Errors

type OverflowException struct {
	message string
	code    int
}

func NewOverflowException(message string, code ...int) Exception {
	code = append(code, 0)
	return OverflowException{
		message: message,
		code:    code[0],
	}
}

func (this OverflowException) GetMessage() string {
	return this.message
}

func (this OverflowException) Error() string {
	return this.GetMessage()
}

func (this OverflowException) GetCode() int {
	return this.code
}
