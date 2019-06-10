package Errors

type Exception interface {
	GetMessage() string
	Error() string
	GetCode() int
}

type exception struct {
	message string
	code    int
}

func NewException(message string, code ...int) Exception {
	code = append(code, 0)
	return exception{
		message: message,
		code:    code[0],
	}
}

func (this exception) GetMessage() string {
	return this.message
}

func (this exception) Error() string {
	return this.GetMessage()
}

func (this exception) GetCode() int {
	return this.code
}
