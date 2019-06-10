package Http

import (
	"context"
	"fmt"
	"github.com/larisgo/framework/Errors"
	"github.com/larisgo/framework/Foundation"
	"github.com/larisgo/framework/Http/HttpFoundation"
	"github.com/larisgo/framework/Support"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type Request struct {
	App *Foundation.Application

	request  *http.Request
	response http.ResponseWriter

	method      string
	pathInfo    string
	isHostValid bool

	Query      *HttpFoundation.ParameterBag
	Post       *HttpFoundation.ParameterBag
	Attributes *HttpFoundation.ParameterBag
	Headers    http.Header

	context       context.Context
	contextCancel context.CancelFunc
}

func NewRequest(app *Foundation.Application, response http.ResponseWriter, request *http.Request) (this *Request) {
	// 解析POST
	this = &Request{App: app, request: request, response: response}
	this.Query = HttpFoundation.NewParameterBag(request.URL.Query())

	request.ParseMultipartForm(64 << 20)
	this.Post = HttpFoundation.NewParameterBag(request.PostForm)
	this.Headers = request.Header
	this.Attributes = HttpFoundation.NewParameterBag(map[string][]string{})

	this.isHostValid = true

	this.context, this.contextCancel = context.WithCancel(request.Context())
	return this
}

// ----- implement context.Context interface ----- //
func (this *Request) Deadline() (time.Time, bool) {
	return this.context.Deadline()
}

func (this *Request) Done() <-chan struct{} {
	return this.context.Done()
}

func (this *Request) Err() error {
	return this.context.Err()
}

func (this *Request) Response() http.ResponseWriter {
	return this.response
}

func (this *Request) Request() *http.Request {
	return this.request
}

/**
 * Returns the path being requested relative to the executed script.
 *
 * The path info always starts with a /.
 *
 * Suppose this request is instantiated from /mysite on localhost:
 *
 *  * http://localhost/mysite              returns an empty string
 *  * http://localhost/mysite/about        returns '/about'
 *  * http://localhost/mysite/enco%20ded   returns '/enco%20ded'
 *  * http://localhost/mysite/about?var=1  returns '/about'
 *
 * @return string The raw path (i.e. not urldecoded)
 */
func (this *Request) GetPathInfo() string {
	if this.pathInfo == "" {
		this.pathInfo = this.request.URL.Path
	}

	return this.pathInfo
}

/**
 * Get the request method.
 *
 * @return string
 */
func (this *Request) Method() string {
	return this.GetMethod()
}

/**
 * Gets a "parameter" value from any bag.
 *
 * This method is mainly useful for libraries that want to provide some flexibility. If you don't need the
 * flexibility in controllers, it is better to explicitly get request parameters from the appropriate
 * public property instead (attributes, query, request).
 *
 * Order of precedence: PATH (routing placeholders or custom attributes), GET, BODY
 *
 * @param string key     The key
 * @param mixed  default The default value if the parameter key does not exist
 *
 * @return string
 */
func (this *Request) Get(key string, _default ...string) string {
	_default = append(_default, "")

	if this.Attributes.Has(key) {
		return this.Attributes.Get(key)
	}

	if this.Query.Has(key) {
		return this.Query.Get(key)
	}

	if this.Post.Has(key) {
		return this.Post.Get(key)
	}

	return _default[0]
}

func (this *Request) Gets(key string, _default ...[]string) []string {
	_default = append(_default, []string{})

	if this.Attributes.Has(key) {
		return this.Attributes.Gets(key)
	}

	if this.Query.Has(key) {
		return this.Query.Gets(key)
	}

	if this.Post.Has(key) {
		return this.Post.Gets(key)
	}

	return _default[0]
}

func (this *Request) GetMethod() string {
	if this.method != "" {
		return this.method
	}

	this.method = strings.ToUpper(this.request.Method)

	if this.method != "POST" {
		return this.method
	}
	method := this.Headers.Get("X-HTTP-METHOD-OVERRIDE")

	if method == "" {
		if this.Post.Has("_method") {
			method = this.Post.Get("_method")
		} else if this.Query.Has("_method") {
			method = this.Query.Get("_method")
		} else {
			method = "POST"
		}
	}

	method = strings.ToUpper(method)
	if _, ok := map[string]bool{"GET": true, "HEAD": true, "POST": true, "PUT": true, "DELETE": true, "CONNECT": true, "OPTIONS": true, "PATCH": true, "PURGE": true, "TRACE": true}[method]; ok {
		this.method = method
		return this.method
	}
	if !regexp.MustCompile(`^[A-Z]+$`).MatchString(method) {
		panic(Errors.NewUnexpectedValueException(fmt.Sprintf(`Invalid method override "%s".`, method)))
	}

	this.method = method

	return this.method
}

func (this *Request) GetHost() string {
	host := this.request.Host
	// trim and remove port number from host
	// host is lowercase as per RFC 952/2181
	host = regexp.MustCompile(`:\d+$`).ReplaceAllString(strings.TrimSpace(host), "")
	// as the host can come from the user (HTTP_HOST and depending on the configuration, SERVER_NAME too can come from the user)
	// check that it does not contain forbidden characters (see RFC 952 and RFC 2181)
	// use preg_replace() instead of preg_match() to prevent DoS attacks with long host names
	if host != "" {
		if host = regexp.MustCompile(`(?:^\[)?[a-zA-Z0-9-:\]_]+\.?`).ReplaceAllString(host, ""); host != "" {
			if !this.isHostValid {
				return ""
			}
			this.isHostValid = false
			panic(Errors.NewUnexpectedValueException(fmt.Sprintf(`Invalid Host "%s".`, host)))
		}
	}

	return host
}

/**
 * Determine if the request is over HTTPS.
 *
 * @return bool
 */
func (this *Request) Secure() bool {
	return this.request.TLS != nil
}

/**
 * Get the current path info for the request.
 *
 * @return string
 */
func (this *Request) Path() string {
	if pattern := strings.Trim(this.GetPathInfo(), "/"); pattern != "" {
		return pattern
	}
	return "/"
}

/**
 * Get the client user agent.
 *
 * @return string
 */
func (this *Request) UserAgent() string {
	return this.Headers.Get("User-Agent")
}

/**
 * Determine if the request is sending JSON.
 *
 * @return bool
 */
func (this *Request) IsJson() bool {
	return Support.Str().Contains([]string{"/json", "+json"}, this.Headers.Get("Content-Type"))
}
