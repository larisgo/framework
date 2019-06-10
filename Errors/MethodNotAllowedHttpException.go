package Errors

import (
	"bytes"
	"strings"
)

type MethodNotAllowedHttpException struct {
	message    string
	code       int
	statusCode int
	headers    map[string][]string
}

func join(v map[string]bool, splite string) string {
	var buf bytes.Buffer
	for k, _ := range v {
		if buf.Len() > 0 {
			buf.WriteString(splite)
		}
		buf.WriteString(strings.ToUpper(k))
	}
	return buf.String()
}

func NewMethodNotAllowedHttpException(allow map[string]bool, message string, code ...int) Exception {
	code = append(code, 0)
	return MethodNotAllowedHttpException{
		message:    message,
		code:       code[0],
		statusCode: 405,
		headers: map[string][]string{
			"Allow": []string{
				join(allow, ", "),
			},
		},
	}
}

func (this MethodNotAllowedHttpException) GetMessage() string {
	return this.message
}

func (this MethodNotAllowedHttpException) Error() string {
	return this.GetMessage()
}

func (this MethodNotAllowedHttpException) GetCode() int {
	return this.code
}

func (this MethodNotAllowedHttpException) GetStatusCode() int {
	return this.statusCode
}

func (this MethodNotAllowedHttpException) GetHeaders() map[string][]string {
	return this.headers
}

/**
 * Set response headers.
 *
 * @param array $headers Response headers
 */
func (this MethodNotAllowedHttpException) SetHeaders(headers map[string][]string) {
	this.headers = headers
}
