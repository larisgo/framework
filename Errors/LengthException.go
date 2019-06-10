package Errors

type LengthException struct {
	message string
	code    int
}

func NewLengthException(message string, code ...int) Exception {
	code = append(code, 0)
	return LengthException{
		message: message,
		code:    code[0],
	}
}

func (this LengthException) GetMessage() string {
	return this.message
}

func (this LengthException) Error() string {
	return this.GetMessage()
}

func (this LengthException) GetCode() int {
	return this.code
}
