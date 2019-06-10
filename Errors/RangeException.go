package Errors

type RangeException struct {
	message string
	code    int
}

func NewRangeException(message string, code ...int) Exception {
	code = append(code, 0)
	return RangeException{
		message: message,
		code:    code[0],
	}
}
func (this RangeException) GetMessage() string {
	return this.message
}

func (this RangeException) Error() string {
	return this.GetMessage()
}

func (this RangeException) GetCode() int {
	return this.code
}
