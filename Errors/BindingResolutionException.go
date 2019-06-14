package Errors

type BindingResolutionException struct {
	message string
	code    int
}

func NewBindingResolutionException(message string, code ...int) Exception {
	code = append(code, 0)
	return BindingResolutionException{
		message: message,
		code:    code[0],
	}
}

func (this BindingResolutionException) GetMessage() string {
	return this.message
}

func (this BindingResolutionException) Error() string {
	return this.GetMessage()
}

func (this BindingResolutionException) GetCode() int {
	return this.code
}
