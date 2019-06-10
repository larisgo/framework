package Http

import (
	"encoding/json"
	"fmt"
	"github.com/larisgo/framework/Errors"
	"github.com/larisgo/framework/Support"
	"net/http"
)

const (
	HTTP_CONTINUE                        = 100
	HTTP_SWITCHING_PROTOCOLS             = 101
	HTTP_PROCESSING                      = 102 // RFC2518
	HTTP_EARLY_HINTS                     = 103 // RFC8297
	HTTP_OK                              = 200
	HTTP_CREATED                         = 201
	HTTP_ACCEPTED                        = 202
	HTTP_NON_AUTHORITATIVE_INFORMATION   = 203
	HTTP_NO_CONTENT                      = 204
	HTTP_RESET_CONTENT                   = 205
	HTTP_PARTIAL_CONTENT                 = 206
	HTTP_MULTI_STATUS                    = 207 // RFC4918
	HTTP_ALREADY_REPORTED                = 208 // RFC5842
	HTTP_IM_USED                         = 226 // RFC3229
	HTTP_MULTIPLE_CHOICES                = 300
	HTTP_MOVED_PERMANENTLY               = 301
	HTTP_FOUND                           = 302
	HTTP_SEE_OTHER                       = 303
	HTTP_NOT_MODIFIED                    = 304
	HTTP_USE_PROXY                       = 305
	HTTP_RESERVED                        = 306
	HTTP_TEMPORARY_REDIRECT              = 307
	HTTP_PERMANENTLY_REDIRECT            = 308 // RFC7238
	HTTP_BAD_REQUEST                     = 400
	HTTP_UNAUTHORIZED                    = 401
	HTTP_PAYMENT_REQUIRED                = 402
	HTTP_FORBIDDEN                       = 403
	HTTP_NOT_FOUND                       = 404
	HTTP_METHOD_NOT_ALLOWED              = 405
	HTTP_NOT_ACCEPTABLE                  = 406
	HTTP_PROXY_AUTHENTICATION_REQUIRED   = 407
	HTTP_REQUEST_TIMEOUT                 = 408
	HTTP_CONFLICT                        = 409
	HTTP_GONE                            = 410
	HTTP_LENGTH_REQUIRED                 = 411
	HTTP_PRECONDITION_FAILED             = 412
	HTTP_REQUEST_ENTITY_TOO_LARGE        = 413
	HTTP_REQUEST_URI_TOO_LONG            = 414
	HTTP_UNSUPPORTED_MEDIA_TYPE          = 415
	HTTP_REQUESTED_RANGE_NOT_SATISFIABLE = 416
	HTTP_EXPECTATION_FAILED              = 417
	HTTP_I_AM_A_TEAPOT                   = 418 // RFC2324
	HTTP_MISDIRECTED_REQUEST             = 421 // RFC7540
	HTTP_UNPROCESSABLE_ENTITY            = 422 // RFC4918
	HTTP_LOCKED                          = 423 // RFC4918
	HTTP_FAILED_DEPENDENCY               = 424 // RFC4918

	/**
	 * @deprecated
	 */
	HTTP_RESERVED_FOR_WEBDAV_ADVANCED_COLLECTIONS_EXPIRED_PROPOSAL = 425 // RFC2817
	HTTP_TOO_EARLY                                                 = 425 // RFC-ietf-httpbis-replay-04
	HTTP_UPGRADE_REQUIRED                                          = 426 // RFC2817
	HTTP_PRECONDITION_REQUIRED                                     = 428 // RFC6585
	HTTP_TOO_MANY_REQUESTS                                         = 429 // RFC6585
	HTTP_REQUEST_HEADER_FIELDS_TOO_LARGE                           = 431 // RFC6585
	HTTP_UNAVAILABLE_FOR_LEGAL_REASONS                             = 451
	HTTP_INTERNAL_SERVER_ERROR                                     = 500
	HTTP_NOT_IMPLEMENTED                                           = 501
	HTTP_BAD_GATEWAY                                               = 502
	HTTP_SERVICE_UNAVAILABLE                                       = 503
	HTTP_GATEWAY_TIMEOUT                                           = 504
	HTTP_VERSION_NOT_SUPPORTED                                     = 505
	HTTP_VARIANT_ALSO_NEGOTIATES_EXPERIMENTAL                      = 506 // RFC2295
	HTTP_INSUFFICIENT_STORAGE                                      = 507 // RFC4918
	HTTP_LOOP_DETECTED                                             = 508 // RFC5842
	HTTP_NOT_EXTENDED                                              = 510 // RFC2774
	HTTP_NETWORK_AUTHENTICATION_REQUIRED                           = 511 // RFC6585
)

type Response struct {
	/**
	 * @var ResponseHeader
	 */
	Headers http.Header

	response http.ResponseWriter

	/**
	 * @var string
	 */
	content []byte

	/**
	 * @var string
	 */
	version string

	/**
	 * @var int
	 */
	statusCode int

	/**
	 * @var string
	 */
	statusText string

	/**
	 * @var string
	 */
	charset string
}

func NewResponse(content interface{}, status int, headers ...map[string][]string) (this *Response) {
	headers = append(headers, map[string][]string{})

	this = &Response{}
	this.Headers = http.Header(headers[0])
	this.SetContent(content)
	this.SetStatusCode(status)
	return this
}

func (this *Response) Prepare(request *Request) *Response {
	this.response = request.Response()

	return this
}

/**
 * Set the content on the response.
 *
 * @param  mixed  $content
 * @return $this
 */
func (this *Response) SetContent(content interface{}) *Response {
	// If the content is "JSONable" we will set the appropriate header and convert
	// the content to JSON. This is useful when returning something like models
	// from routes that will be automatically transformed to their JSON form.
	if this.shouldBeJson(content) {
		this.Header("Content-Type", "application/json")

		if data, err := this.morphToJson(content); err != nil {
			panic(err)
		} else {
			this.content = data
		}
	} else if _string, ok := content.(string); ok {
		this.content = []byte(_string)
	} else if _byte, ok := content.([]byte); ok {
		this.content = _byte
	} else {
		panic(Errors.NewUnexpectedValueException(`The Response content must be a string or []byte`))
	}

	// this.content = content

	return this
}

/**
 * Determine if the given content should be turned into JSON.
 *
 * @param  mixed  $content
 * @return bool
 */
func (this *Response) shouldBeJson(content interface{}) bool {
	if _, ok := content.(Support.Jsonable); ok {
		return ok
	}
	if _, ok := content.(map[string]interface{}); ok {
		return ok
	}
	return false
}

/**
 * Morph the given content into JSON.
 *
 * @param  mixed   $content
 * @return string
 */
func (this *Response) morphToJson(content interface{}) ([]byte, error) {
	if _jsonable, ok := content.(Support.Jsonable); ok {
		return _jsonable.ToJson()
	}
	return json.Marshal(content)
}

/**
 * Sets the response status code.
 *
 * If the status text is null it will be automatically populated for the known
 * status codes and left empty otherwise.
 *
 * @return $this
 *
 * @throws \InvalidArgumentException When the HTTP status code is not valid
 *
 * @final
 */
func (this *Response) SetStatusCode(code int) *Response {
	this.statusCode = code
	if this.IsInvalid() {
		panic(Errors.NewInvalidArgumentException(fmt.Sprintf(`The HTTP status code "%d" is not valid.`, code)))
	}

	return this
}

/**
 * Set a header on the Response.
 *
 * @param  string  $key
 * @param  array|string  $values
 * @param  bool    $replace
 * @return $this
 */
func (this *Response) Header(key string, value string, replace ...bool) *Response {
	replace = append(replace, true)
	if replace[0] {
		this.Headers.Set(key, value)
	} else {
		this.Headers.Add(key, value)
	}

	return this
}

/**
 * Get the status code for the response.
 *
 * @return int
 */
func (this *Response) Status() int {
	return this.statusCode
}

/**
 * Get the content of the response.
 *
 * @return string
 */
func (this *Response) Content() []byte {
	return this.content
}

/**
 * Get the content of the response.
 *
 * @return string
 */
func (this *Response) ContentString() string {
	return string(this.content)
}

/**
 * Sends HTTP headers.
 *
 * @return $this
 */
func (this *Response) SendHeaders() *Response {
	for key, values := range this.Headers {
		for _, value := range values {
			this.response.Header().Add(key, value)
		}
	}
	return this
}

/**
 * Sends content for the current web response.
 *
 * @return $this
 */
func (this *Response) SendContent() *Response {
	this.response.WriteHeader(this.statusCode)
	this.response.Write(this.content)

	return this
}

func (this *Response) Send() *Response {
	this.SendHeaders()
	this.SendContent()

	return this
}

/**
 * Is response invalid?
 *
 * @see http://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html
 *
 * @final
 */
func (this *Response) IsInvalid() bool {
	return this.statusCode < 100 || this.statusCode >= 600
}

/**
 * Is response informative?
 *
 * @final
 */
func (this *Response) IsInformational() bool {
	return this.statusCode >= 100 && this.statusCode < 200
}

/**
 * Is response successful?
 *
 * @final
 */
func (this *Response) IsSuccessful() bool {
	return this.statusCode >= 200 && this.statusCode < 300
}

/**
 * Is the response a redirect?
 *
 * @final
 */
func (this *Response) IsRedirection() bool {
	return this.statusCode >= 300 && this.statusCode < 400
}

/**
 * Is there a client error?
 *
 * @final
 */
func (this *Response) IsClientError() bool {
	return this.statusCode >= 400 && this.statusCode < 500
}

/**
 * Was there a server side error?
 *
 * @final
 */
func (this *Response) IsServerError() bool {
	return this.statusCode >= 500 && this.statusCode < 600
}

/**
 * Is the response OK?
 *
 * @final
 */
func (this *Response) IsOk() bool {
	return 200 == this.statusCode
}

/**
 * Is the response forbidden?
 *
 * @final
 */
func (this *Response) IsForbidden() bool {
	return 403 == this.statusCode
}

/**
 * Is the response a not found error?
 *
 * @final
 */
func (this *Response) IsNotFound() bool {
	return 404 == this.statusCode
}

/**
 * Is the response a redirect of some form?
 *
 * @final
 */
func (this *Response) IsRedirect(location ...string) bool {
	location = append(location, "")
	_, ok := map[int]bool{
		201: true, 301: true, 302: true, 303: true, 307: true, 308: true,
	}[this.statusCode]

	return ok && ((location[0] == "") || (location[0] != "") && (location[0] == this.Headers.Get("Location")))
}

/**
 * Is the response empty?
 *
 * @final
 */
func (this *Response) IsEmpty() bool {
	_, ok := map[int]bool{
		204: true, 304: true,
	}[this.statusCode]
	return ok
}
