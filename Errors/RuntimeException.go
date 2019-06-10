package Errors

type RuntimeException struct {
	message string
	code    int
}

func NewRuntimeException(message string, code ...int) Exception {
	code = append(code, 0)
	return RuntimeException{
		message: message,
		code:    code[0],
	}
}
func (this RuntimeException) GetMessage() string {
	return this.message
}

func (this RuntimeException) Error() string {
	return this.GetMessage()
}

func (this RuntimeException) GetCode() int {
	return this.code
}
