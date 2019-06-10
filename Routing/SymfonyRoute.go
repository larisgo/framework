package Routing

import (
	"fmt"
	"github.com/larisgo/framework/Errors"
	"regexp"
	"strings"
)

type SymfonyRoute struct {
	path         string
	host         string
	schemes      map[string]bool
	methods      map[string]bool
	defaults     map[string]string
	requirements map[string]string
	options      map[string]interface{}
	condition    string

	compiled *CompiledRoute
}

func NewSymfonyRoute(path string, defaults map[string]string, requirements map[string]string, options map[string]interface{}, host string, schemes map[string]bool, methods map[string]bool, condition string) (this *SymfonyRoute) {
	this = &SymfonyRoute{}

	this.schemes = map[string]bool{}
	this.methods = map[string]bool{}
	this.defaults = map[string]string{}
	this.requirements = map[string]string{}
	this.options = map[string]interface{}{}

	this.SetPath(path)
	this.AddDefaults(defaults)
	this.AddRequirements(requirements)
	this.SetOptions(options)
	this.SetHost(host)
	this.SetSchemes(schemes)
	this.SetMethods(methods)
	this.SetCondition(condition)

	return this
}

/**
 * Returns the pattern for the path.
 *
 * @return string The path pattern
 */
func (this *SymfonyRoute) GetPath() string {
	return this.path
}

/**
 * Sets the pattern for the path.
 *
 * This method implements a fluent interface.
 *
 * @param string pattern The path pattern
 *
 * @return this
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
 * Returns the pattern for the host.
 *
 * @return string The host pattern
 */
func (this *SymfonyRoute) GetHost() string {
	return this.host
}

/**
 * Sets the pattern for the host.
 *
 * This method implements a fluent interface.
 *
 * @param string pattern The host pattern
 *
 * @return this
 */
func (this *SymfonyRoute) SetHost(pattern string) *SymfonyRoute {
	this.host = pattern
	this.compiled = nil

	return this
}

/**
 * Returns the lowercased schemes this route is restricted to.
 * So an empty array means that any scheme is allowed.
 *
 * @return string[] The schemes
 */
func (this *SymfonyRoute) GetSchemes() map[string]bool {
	return this.schemes
}

/**
 * Sets the schemes (e.g. 'https') this route is restricted to.
 * So an empty array means that any scheme is allowed.
 *
 * This method implements a fluent interface.
 *
 * @param map[string]bool schemes The scheme or an array of schemes
 *
 * @return this
 */
func (this *SymfonyRoute) SetSchemes(schemes map[string]bool) *SymfonyRoute {
	var _schemes map[string]bool
	for k, v := range schemes {
		_schemes[strings.ToUpper(k)] = v
	}
	this.schemes = _schemes
	this.compiled = nil

	return this
}

/**
 * Checks if a scheme requirement has been set.
 *
 * @param string scheme
 *
 * @return bool true if the scheme requirement exists, otherwise false
 */
func (this *SymfonyRoute) HasScheme(scheme string) bool {
	v, ok := this.schemes[scheme]
	return ok && v
}

/**
 * Returns the uppercased HTTP methods this route is restricted to.
 * So an empty array means that any method is allowed.
 *
 * @return string[] The methods
 */
func (this *SymfonyRoute) GetMethods() map[string]bool {
	return this.methods
}

/**
 * Sets the HTTP methods (e.g. 'POST') this route is restricted to.
 * So an empty array means that any method is allowed.
 *
 * This method implements a fluent interface.
 *
 * @param string|string[] $methods The method or an array of methods
 *
 * @return $this
 */
func (this *SymfonyRoute) SetMethods(methods map[string]bool) *SymfonyRoute {
	var _methods map[string]bool
	for k, v := range methods {
		_methods[strings.ToUpper(k)] = v
	}
	this.methods = _methods
	this.compiled = nil

	return this
}

/**
 * Returns the options.
 *
 * @return array The options
 */
func (this *SymfonyRoute) GetOptions() map[string]interface{} {
	return this.options
}

/**
 * Sets the options.
 *
 * This method implements a fluent interface.
 *
 * @param array options The options
 *
 * @return this
 */
func (this *SymfonyRoute) SetOptions(options map[string]interface{}) *SymfonyRoute {
	this.options = map[string]interface{}{
		`compiler_class`: NewSymfonyRouteCompiler,
	}
	return this.AddOptions(options)
}

/**
 * Adds options.
 *
 * This method implements a fluent interface.
 *
 * @param array options The options
 *
 * @return this
 */
func (this *SymfonyRoute) AddOptions(options map[string]interface{}) *SymfonyRoute {
	for name, option := range options {
		this.options[name] = option
	}
	this.compiled = nil

	return this
}

/**
 * Sets an option value.
 *
 * This method implements a fluent interface.
 *
 * @param string name  An option name
 * @param mixed  value The option value
 *
 * @return this
 */
func (this *SymfonyRoute) SetOption(name string, value interface{}) *SymfonyRoute {
	this.options[name] = value
	this.compiled = nil

	return this
}

/**
 * Get an option value.
 *
 * @param string $name An option name
 *
 * @return mixed The option value or null when not given
 */
func (this *SymfonyRoute) GetOption(name string) interface{} {
	if v, ok := this.options[name]; ok {
		return v
	}
	return nil
}

/**
 * Checks if an option has been set.
 *
 * @param string $name An option name
 *
 * @return bool true if the option is set, false otherwise
 */
func (this *SymfonyRoute) HasOption(name string) bool {
	_, ok := this.options[name]
	return ok
}

/**
 * Returns the defaults.
 *
 * @return array The defaults
 */
func (this *SymfonyRoute) GetDefaults() map[string]string {
	return this.defaults
}

/**
 * Sets the defaults.
 *
 * This method implements a fluent interface.
 *
 * @param array $defaults The defaults
 *
 * @return this
 */
func (this *SymfonyRoute) SetDefaults(defaults map[string]string) *SymfonyRoute {
	this.defaults = map[string]string{}

	return this.AddDefaults(defaults)
}

/**
 * Adds defaults.
 *
 * This method implements a fluent interface.
 *
 * @param array $defaults The defaults
 *
 * @return this
 */
func (this *SymfonyRoute) AddDefaults(defaults map[string]string) *SymfonyRoute {
	for name, _default := range defaults {
		this.defaults[name] = _default
	}
	this.compiled = nil

	return this
}

/**
 * Gets a default value.
 *
 * @param string $name A variable name
 *
 * @return mixed The default value or null when not given
 */
func (this *SymfonyRoute) GetDefault(name string) string {
	if v, ok := this.defaults[name]; ok {
		return v
	}
	return ""
}

/**
 * Checks if a default value is set for the given variable.
 *
 * @param string $name A variable name
 *
 * @return bool true if the default value is set, false otherwise
 */
func (this *SymfonyRoute) HasDefault(name string) bool {
	_, ok := this.defaults[name]
	return ok
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
 * Returns the requirements.
 *
 * @return array The requirements
 */
func (this *SymfonyRoute) GetRequirements() map[string]string {
	return this.requirements
}

/**
 * Sets the requirements.
 *
 * This method implements a fluent interface.
 *
 * @param array $requirements The requirements
 *
 * @return this
 */
func (this *SymfonyRoute) SetRequirements(requirements map[string]string) *SymfonyRoute {
	this.requirements = map[string]string{}

	return this.AddRequirements(requirements)
}

/**
 * Adds requirements.
 *
 * This method implements a fluent interface.
 *
 * @param array requirements The requirements
 *
 * @return this
 */
func (this *SymfonyRoute) AddRequirements(requirements map[string]string) *SymfonyRoute {
	for key, regex := range requirements {
		this.requirements[key] = this.sanitizeRequirement(key, regex)
	}
	this.compiled = nil

	return this
}

/**
 * Returns the requirement for the given key.
 *
 * @param string $key The key
 *
 * @return string The regex or null when not given
 */
func (this *SymfonyRoute) GetRequirement(key string) string {
	if v, ok := this.requirements[key]; ok {
		return v
	}
	return ""
}

/**
 * Checks if a requirement is set for the given key.
 *
 * @param string $key A variable name
 *
 * @return bool true if a requirement is specified, false otherwise
 */
func (this *SymfonyRoute) HasRequirement(key string) bool {
	_, ok := this.requirements[key]
	return ok
}

/**
 * Sets a requirement for the given key.
 *
 * @param string key   The key
 * @param string regex The regex
 *
 * @return this
 */
func (this *SymfonyRoute) SetRequirement(key, regex string) *SymfonyRoute {
	this.requirements[key] = this.sanitizeRequirement(key, regex)
	this.compiled = nil

	return this
}

/**
 * Returns the condition.
 *
 * @return string The condition
 */
func (this *SymfonyRoute) GetCondition() string {
	return this.condition
}

/**
 * Sets the condition.
 *
 * This method implements a fluent interface.
 *
 * @param string $condition The condition
 *
 * @return this
 */
func (this *SymfonyRoute) SetCondition(condition string) *SymfonyRoute {
	this.condition = condition
	this.compiled = nil

	return this
}

/**
 * Compiles the route.
 *
 * @return CompiledRoute A CompiledRoute instance
 *
 * @throws \LogicException If the Route cannot be compiled because the
 *                         path or host pattern is invalid
 *
 * @see RouteCompiler which is responsible for the compilation process
 */
func (this *SymfonyRoute) Compile() *CompiledRoute {
	if this.compiled != nil {
		return this.compiled
	}
	class := this.GetOption("compiler_class")
	this.compiled = class.(func() *SymfonyRouteCompiler)().Compile(this)
	return this.compiled
}

func (this *SymfonyRoute) sanitizeRequirement(key string, regex string) string {

	if "" != regex && "^" == regex[0:1] {
		regex = regex[1:] // returns false for a single character
	}

	if regexl := len(regex); "" != regex && "$" == regex[regexl-1:] {
		regex = regex[0 : regexl-1]
	}

	if "" == regex {
		panic(Errors.NewInvalidArgumentException(fmt.Sprintf(`Routing requirement for "%s" cannot be empty.`, key)))
	}

	return regex
}
