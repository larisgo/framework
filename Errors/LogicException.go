package Errors

type LogicException struct {
	message string
	code    int
}

func NewLogicException(message string, code ...int) Exception {
	code = append(code, 0)
	return LogicException{
		message: message,
		code:    code[0],
	}
}

func (this LogicException) GetMessage() string {
	return this.message
}

func (this LogicException) Error() string {
	return this.GetMessage()
}

func (this LogicException) GetCode() int {
	return this.code
}
