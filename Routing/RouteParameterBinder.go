package Routing

import (
	"github.com/larisgo/framework/Http"
	"regexp"
	"strings"
)

type RouteParameterBinder struct {
	route *Route
}

func NewRouteParameterBinder(route *Route) *RouteParameterBinder {
	return &RouteParameterBinder{route: route}
}

/**
 * Get the parameters for the route.
 *
 * @param  Http.Request  request
 * @return map[string]string
 */
func (this *RouteParameterBinder) Parameters(request *Http.Request) map[string]string {
	// If the route has a regular expression for the host part of the URI, we will
	// compile that and get the parameter matches for this domain. We will then
	// merge them into this parameters array so that this array is completed.
	parameters := this.bindPathParameters(request)

	// If the route has a regular expression for the host part of the URI, we will
	// compile that and get the parameter matches for this domain. We will then
	// merge them into this parameters array so that this array is completed.
	if this.route.compiled.GetHostRegex() != "" {
		parameters = this.bindHostParameters(request, parameters)
	}

	return this.replaceDefaults(parameters)
}

func (this *RouteParameterBinder) combine(key []string, value []string) map[string]string {
	result := map[string]string{}
	if len(key) <= len(value) {
		for k, v := range key {
			result[v] = value[k]
		}
	}
	return result
}

/**
 * Get the parameter matches for the path portion of the URI.
 *
 * @param  Http.Request  request
 * @return array
 */
func (this *RouteParameterBinder) bindPathParameters(request *Http.Request) map[string]string {
	path := "/" + strings.TrimLeft(request.Path(), "/")
	_regexp := regexp.MustCompile(this.route.compiled.GetRegex())
	result := this.combine(_regexp.SubexpNames()[1:], _regexp.FindStringSubmatch(path)[1:])

	return this.matchToKeys(result)
}

/**
 * Extract the parameter list from the host part of the request.
 *
 * @param  Http.Request  request
 * @param  array  parameters
 * @return array
 */
func (this *RouteParameterBinder) bindHostParameters(request *Http.Request, parameters map[string]string) map[string]string {
	_regexp := regexp.MustCompile(this.route.compiled.GetHostRegex())
	result := this.combine(_regexp.SubexpNames()[1:], _regexp.FindStringSubmatch(request.GetHost())[1:])

	for k, v := range this.matchToKeys(result) {
		parameters[k] = v
	}

	return parameters
}

/**
 * Combine a set of parameter matches with the route's keys.
 *
 * @param  array  matches
 * @return array
 */
func (this *RouteParameterBinder) matchToKeys(matches map[string]string) map[string]string {
	parameterNames := this.route.ParameterNames()
	if parameterNames == nil {
		return map[string]string{}
	}
	for k, _ := range matches {
		if _, ok := parameterNames[k]; !ok {
			delete(matches, k)
		}
	}

	return matches
}

/**
 * Replace null parameters with their defaults.
 *
 * @param  array  parameters
 * @return array
 */
func (this *RouteParameterBinder) replaceDefaults(parameters map[string]string) map[string]string {
	for key, value := range parameters {
		if value == "" {
			if v, ok := this.route.defaults[key]; ok {
				parameters[key] = v
			}
		}
	}

	for key, value := range this.route.defaults {
		if _, ok := parameters[key]; !ok {
			parameters[key] = value
		}
	}

	return parameters
}
