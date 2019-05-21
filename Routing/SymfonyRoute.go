package Routing

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type SymfonyRoute struct {
	path         string
	host         string
	schemes      map[string]bool
	methods      []string
	defaults     map[string]string
	requirements map[string]string
	options      map[string]interface{}
	condition    string

	compiled *CompiledRoute
}

func NewSymfonyRoute(path string, defaults map[string]string, requirements map[string]string, options map[string]interface{}, host string, schemes map[string]bool, methods []string, condition string) (this *SymfonyRoute) {
	this = &SymfonyRoute{}

	this.schemes = make(map[string]bool)
	this.methods = make([]string, 0)
	this.defaults = make(map[string]string)
	this.requirements = make(map[string]string)
	this.options = make(map[string]interface{})

	this.SetPath(path)
	// this.AddDefaults(defaults)
	// this.AddRequirements(requirements)
	// this.SetOptions(options)
	// this.SetHost(host)
	// this.SetSchemes(schemes)
	// this.SetMethods(methods)
	// this.SetCondition(condition)

	return this
}

/**
 * Sets the pattern for the path.
 *
 * This method implements a fluent interface.
 *
 * @param string pattern The path pattern
 *
 * @return $this
 */
func (this *SymfonyRoute) SetPath(pattern string) *SymfonyRoute {
	if strings.IndexAny(pattern, "?<") > -1 {
		regex := regexp.MustCompile(`\{(\w+)(<.*?>)?(\?[^\}]*)?\}`)
		pattern = regex.ReplaceAllStringFunc(pattern, func(m string) string {
			_m := regex.FindStringSubmatch(m)
			if len(_m[3]) > 0 {
				if _m[3] != "?" {
					_m[3] = _m[3][1:]
				} else {
					_m[3] = ""
				}
				this.SetDefault(_m[1], _m[3])
			}
			if _m2l := len(_m[2]); _m2l > 0 {
				_m[2] = _m[2][1 : _m2l-1]
				this.SetRequirement(_m[1], _m[2])
			}
			return "{" + _m[1] + "}"
		})
	}

	// A pattern must start with a slash and must not have multiple slashes at the beginning because the
	// generated path for this route would be confused with a network path, e.g. '//domain.com/path'.
	this.path = "/" + strings.TrimLeft(strings.TrimSpace(pattern), "/")
	this.compiled = nil

	return this
}

/**
 * Sets a default value.
 *
 * @param string name    A variable name
 * @param mixed  default The default value
 *
 * @return this
 */
func (this *SymfonyRoute) SetDefault(name, _default string) *SymfonyRoute {
	this.defaults[name] = _default
	this.compiled = nil

	return this
}

/**
 * Sets a requirement for the given key.
 *
 * @param string $key   The key
 * @param string $regex The regex
 *
 * @return $this
 */
func (this *SymfonyRoute) SetRequirement(key, regex string) *SymfonyRoute {
	this.requirements[key] = this.sanitizeRequirement(key, regex)
	this.compiled = nil

	return this
}

func (this *SymfonyRoute) sanitizeRequirement(key, regex string) string {

	if "" != regex && "^" == regex[0:1] {
		regex = regex[1:] // returns false for a single character
	}

	if regexl := len(regex); "" != regex && "$" == regex[regexl-1:] {
		regex = regex[0 : regexl-1]
	}

	if "" == regex {
		panic(errors.New(fmt.Sprintf(`Routing requirement for "%s" cannot be empty.`, key)))
	}

	return regex
}
