package Routing

import (
	"strings"
)

type RouteGroup struct {
}

func NewRouteGroup() *RouteGroup {
	return &RouteGroup{}
}

func (this *RouteGroup) Merge(_new map[string]string, _old map[string]string) map[string]string {
	_ntmp := this.clone(_new)
	_otmp := this.clone(_old)
	if _, ok := _ntmp["domain"]; ok {
		delete(_otmp, "domain")
	}
	_ntmp = this.formatAs(_ntmp, _otmp)
	_ntmp["prefix"] = this.formatPrefix(_ntmp, _otmp)

	delete(_otmp, "prefix")
	delete(_otmp, "as")
	for k, v := range _ntmp {
		_otmp[k] = v
	}
	return _otmp
}

func (this *RouteGroup) clone(value map[string]string) (data map[string]string) {
	data = map[string]string{}
	for k, v := range value {
		data[k] = v
	}
	return data
}

/**
 * Format the prefix for the new group attributes.
 *
 * @param  array  new
 * @param  array  old
 * @return string|null
 */
func (this *RouteGroup) formatPrefix(_new map[string]string, _old map[string]string) string {
	_prefix := ""
	if v, ok := _old["prefix"]; ok {
		_prefix = v
	}
	if _, ok := _new["prefix"]; ok {
		return strings.Trim(_prefix, "/") + "/" + strings.Trim(_new["prefix"], "/")
	} else {
		return _prefix
	}
}

/**
 * Format the "as" clause of the new group attributes.
 *
 * @param  array  new
 * @param  array  old
 * @return array
 */
func (this *RouteGroup) formatAs(_new map[string]string, _old map[string]string) map[string]string {
	if ov, ook := _old["as"]; ook {
		if _, nok := _new["as"]; nok {
			_new["as"] = ov + _new["as"]
		} else {
			_new["as"] = ov
		}
	}

	return _new
}
