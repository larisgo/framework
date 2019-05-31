package Http

import (
	"mime/multipart"
)

const HEADER_FORWARDED = 0x1 // When using RFC 7239
const HEADER_X_FORWARDED_FOR = 0x2
const HEADER_X_FORWARDED_HOST = 0x4
const HEADER_X_FORWARDED_PROTO = 0x8
const HEADER_X_FORWARDED_PORT = 0x10
const HEADER_X_FORWARDED_ALL = 0x1e     // All "X-Forwarded-*" headers
const HEADER_X_FORWARDED_AWS_ELB = 0x1a // AWS ELB doesn't send X-Forwarded-Host

const METHOD_HEAD = "HEAD"
const METHOD_GET = "GET"
const METHOD_POST = "POST"
const METHOD_PUT = "PUT"
const METHOD_PATCH = "PATCH"
const METHOD_DELETE = "DELETE"
const METHOD_PURGE = "PURGE"
const METHOD_OPTIONS = "OPTIONS"
const METHOD_TRACE = "TRACE"
const METHOD_CONNECT = "CONNECT"

/**
 * @var string[]
 */
var trustedProxies []string

/**
 * @var string[]
 */
var trustedHostPatterns []string

/**
 * @var string[]
 */
var trustedHosts []string

var httpMethodParameterOverride bool = false

/**
 * @var array
 */
var formats []string

// var requestFactory;

var trustedHeaderSet int = -1

var forwardedParams map[int]string = map[int]string{
	HEADER_X_FORWARDED_FOR:   "for",
	HEADER_X_FORWARDED_HOST:  "host",
	HEADER_X_FORWARDED_PROTO: "proto",
	HEADER_X_FORWARDED_PORT:  "host",
}

/**
 * Names for headers that can be trusted when
 * using trusted proxies.
 *
 * The FORWARDED header is the standard as of rfc7239.
 *
 * The other headers are non-standard, but widely used
 * by popular reverse proxies (like Apache mod_proxy or Amazon EC2).
 */
var trustedHeaders map[int]string = map[int]string{
	HEADER_FORWARDED:         "FORWARDED",
	HEADER_X_FORWARDED_FOR:   "X_FORWARDED_FOR",
	HEADER_X_FORWARDED_HOST:  "X_FORWARDED_HOST",
	HEADER_X_FORWARDED_PROTO: "X_FORWARDED_PROTO",
	HEADER_X_FORWARDED_PORT:  "X_FORWARDED_PORT",
}

type SymfonyRequest struct {
	Attributes             map[string][]string
	Request                map[string][]string
	Query                  map[string][]string
	Server                 map[string][]string
	Files                  map[string][]*multipart.FileHeader
	Cookies                map[string][]string
	Headers                map[string][]string
	content                string
	languages              []string
	charsets               []string
	encodings              []string
	acceptableContentTypes []string
	pathInfo               string
	requestUri             string
	baseUrl                string
	basePath               string
	method                 string
	format                 string
	session                interface{}
	locale                 string
	defaultLocale          string
	isHostValid            bool
	isForwardedValid       bool
}

/**
 * @param array                query      The GET parameters
 * @param array                request    The POST parameters
 * @param array                attributes The request attributes (parameters parsed from the PATH_INFO, ...)
 * @param array                cookies    The COOKIE parameters
 * @param array                files      The FILES parameters
 * @param array                server     The SERVER parameters
 * @param string|resource|null content    The raw body data
 */
func NewSymfonyRequest(query map[string][]string, request map[string][]string, attributes map[string][]string, cookies map[string][]string, files map[string][]*multipart.FileHeader, server map[string][]string, content string) (this *SymfonyRequest) {
	this = &SymfonyRequest{}
	this.defaultLocale = "en"
	this.isHostValid = true
	this.isForwardedValid = true
	this.Initialize(query, request, attributes, cookies, files, server, content)
	return this
}

/**
 * Sets the parameters for this request.
 *
 * This method also re-initializes all properties.
 *
 * @param array                query      The GET parameters
 * @param array                request    The POST parameters
 * @param array                attributes The request attributes (parameters parsed from the PATH_INFO, ...)
 * @param array                cookies    The COOKIE parameters
 * @param array                files      The FILES parameters
 * @param array                server     The SERVER parameters
 * @param string|resource|null content    The raw body data
 */
func (this *SymfonyRequest) Initialize(query map[string][]string, request map[string][]string, attributes map[string][]string, cookies map[string][]string, files map[string][]*multipart.FileHeader, server map[string][]string, content string) {
	// this.request = NewParameterBag(request)
	// this.query = NewParameterBag(query)
	// this.attributes = NewParameterBag(attributes)
	// this.cookies = NewParameterBag(cookies)
	// this.files = NewFileBag(files)
	// this.server = NewServerBag(server)
	// this.headers = NewHeaderBag(this.server.GetHeaders())

	this.content = content
	this.languages = nil
	this.charsets = nil
	this.encodings = nil
	this.acceptableContentTypes = nil
	this.pathInfo = ""
	this.requestUri = ""
	this.baseUrl = ""
	this.basePath = ""
	this.method = ""
	this.format = ""
}
