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
	if _, ok := _new["domain"]; ok {
		delete(_old, "domain")
	}
	_new = this.formatAs(_new, _old)
	_new["prefix"] = this.formatPrefix(_new, _old)

	delete(_old, "prefix")
	delete(_old, "as")
	for k, v := range _new {
		_old[k] = v
	}
	return _old
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
