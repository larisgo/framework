package Errors

type BadMethodCallException struct {
	message string
	code    int
}

func NewBadMethodCallException(message string, code ...int) Exception {
	code = append(code, 0)
	return BadMethodCallException{
		message: message,
		code:    code[0],
	}
}

func (this BadMethodCallException) GetMessage() string {
	return this.message
}

func (this BadMethodCallException) Error() string {
	return this.GetMessage()
}

func (this BadMethodCallException) GetCode() int {
	return this.code
}
