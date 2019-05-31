package Http

import (
	"net/http"
	"strings"
)

type Request struct {
	method  string
	Request *http.Request
}

// func Capture()
// {
//     enableHttpMethodParameterOverride();

//     return createFromBase(SymfonyRequest::createFromGlobals());
// }

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
// public function get($key, $default = null)
// {
//     if ($this !== $result = $this->attributes->get($key, $this)) {
//         return $result;
//     }

//     if ($this !== $result = $this->query->get($key, $this)) {
//         return $result;
//     }

//     if ($this !== $result = $this->request->get($key, $this)) {
//         return $result;
//     }

//     return $default;
// }

func (this *Request) GetMethod() string {
	if this.method != "" {
		return this.method
	}

	this.method = strings.ToUpper(this.Request.Method)

	if this.method != "POST" {
		return this.method
	}

	method := this.Request.Header.Get("X-HTTP-METHOD-OVERRIDE")

	if method == "" && httpMethodParameterOverride {
		// method = this.Post.Get("_method", $this->query->get('_method', 'POST'));
	}

	// if (!$method && self::$httpMethodParameterOverride) {
	//     $method = $this->request->get('_method', $this->query->get('_method', 'POST'));
	// }

	// if (!\is_string($method)) {
	//     return $this->method;
	// }

	// $method = strtoupper($method);

	// if (\in_array($method, ['GET', 'HEAD', 'POST', 'PUT', 'DELETE', 'CONNECT', 'OPTIONS', 'PATCH', 'PURGE', 'TRACE'], true)) {
	//     return $this->method = $method;
	// }

	// if (!preg_match('/^[A-Z]++$/D', $method)) {
	//     throw new SuspiciousOperationException(sprintf('Invalid method override "%s".', $method));
	// }
	// this.Method = method
	return this.method
}
