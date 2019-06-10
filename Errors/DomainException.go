package Errors

type DomainException struct {
	message string
	code    int
}

func NewDomainException(message string, code ...int) Exception {
	code = append(code, 0)
	return DomainException{
		message: message,
		code:    code[0],
	}
}

func (this DomainException) GetMessage() string {
	return this.message
}

func (this DomainException) Error() string {
	return this.GetMessage()
}

func (this DomainException) GetCode() int {
	return this.code
}
