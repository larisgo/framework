package Errors

type UnderflowException struct {
	message string
	code    int
}

func NewUnderflowException(message string, code ...int) Exception {
	code = append(code, 0)
	return UnderflowException{
		message: message,
		code:    code[0],
	}
}
func (this UnderflowException) GetMessage() string {
	return this.message
}

func (this UnderflowException) Error() string {
	return this.GetMessage()
}

func (this UnderflowException) GetCode() int {
	return this.code
}
