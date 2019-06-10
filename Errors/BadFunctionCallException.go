package Errors

type BadFunctionCallException struct {
	message string
	code    int
}

func NewBadFunctionCallException(message string, code ...int) Exception {
	code = append(code, 0)
	return BadFunctionCallException{
		message: message,
		code:    code[0],
	}
}

func (this BadFunctionCallException) GetMessage() string {
	return this.message
}

func (this BadFunctionCallException) Error() string {
	return this.GetMessage()
}

func (this BadFunctionCallException) GetCode() int {
	return this.code
}
