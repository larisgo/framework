package Errors

type NotFoundHttpException struct {
	message    string
	code       int
	statusCode int
	headers    map[string][]string
}

func NewNotFoundHttpException(message string, code ...int) Exception {
	code = append(code, 0)
	return NotFoundHttpException{
		message:    message,
		code:       code[0],
		statusCode: 404,
		headers:    map[string][]string{},
	}
}

func (this NotFoundHttpException) GetMessage() string {
	return this.message
}

func (this NotFoundHttpException) Error() string {
	return this.GetMessage()
}

func (this NotFoundHttpException) GetCode() int {
	return this.code
}

func (this NotFoundHttpException) GetStatusCode() int {
	return this.statusCode
}

func (this NotFoundHttpException) GetHeaders() map[string][]string {
	return this.headers
}

/**
 * Set response headers.
 *
 * @param array $headers Response headers
 */
func (this NotFoundHttpException) SetHeaders(headers map[string][]string) {
	this.headers = headers
}
