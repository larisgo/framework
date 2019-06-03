package Http

import (
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"regexp"
	"strings"
)

type Request struct {
	method        string
	isHostValid   bool
	routeResolver func() interface{}
	Context       *fasthttp.RequestCtx
}

func Capture(request *fasthttp.RequestCtx) (this *Request) {
	this = &Request{Context: request}
	this.isHostValid = true
	this.isHostValid = true
	return this
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
 * @param string $key     The key
 * @param mixed  $default The default value if the parameter key does not exist
 *
 * @return mixed
 */
func (this *Request) Get(key string, _default string) string {
	// if ($this !== $result = $this->attributes->get($key, $this)) {
	//     return $result;
	// }

	if this.Context.QueryArgs().Has(key) {
		return string(this.Context.QueryArgs().Peek(key))
	}

	if this.Context.PostArgs().Has(key) {
		return string(this.Context.PostArgs().Peek(key))
	}

	return _default
}

func (this *Request) GetMethod() string {
	if this.method != "" {
		return this.method
	}

	this.method = strings.ToUpper(string(this.Context.Method()))

	if !this.Context.IsPost() {
		return this.method
	}
	method := string(this.Context.Request.Header.Peek("X-HTTP-METHOD-OVERRIDE"))

	if method == "" {
		if this.Context.PostArgs().Has("_method") {
			method = string(this.Context.PostArgs().Peek("_method"))
		} else if this.Context.QueryArgs().Has("_method") {
			method = string(this.Context.QueryArgs().Peek("_method"))
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
		panic(errors.New(fmt.Sprintf(`Invalid method override "%s".`, method)))
	}
	this.method = method
	return this.method
}

func (this *Request) GetHost() string {
	host := string(this.Context.Host())
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
			panic(errors.New(fmt.Sprintf(`Invalid Host "%s".`, host)))
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
	return this.Context.IsTLS()
}

/**
 * Get the current path info for the request.
 *
 * @return string
 */
func (this *Request) Path() string {
	if pattern := strings.Trim(string(this.Context.Path()), "/"); pattern != "" {
		return pattern
	}
	return "/"
}

// /**
//  * Get the route resolver callback.
//  *
//  * @return \Closure
//  */
// func (this *Request) GetRouteResolver() func() interface{} {
// 	if this.routeResolver != nil {
// 		return this.routeResolver
// 	}
// 	return func() interface{} {
// 		return nil
// 	}
// }

// /**
//  * Set the route resolver callback.
//  *
//  * @param  \Closure  $callback
//  * @return $this
//  */
// func (this *Request) SetRouteResolver(callback func() interface{}) *Request {
// 	this.routeResolver = callback

// 	return this
// }
