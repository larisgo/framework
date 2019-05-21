package Support

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

type str struct {
	_str        string
	snakeCache  []string
	camelCache  []string
	studlyCache []string
}

func Str(_str ...string) *str {
	_str = append(_str, "")
	return &str{_str: _str[0]}
}

/**
 * Determine if a given string matches a given pattern.
 *
 * @param  []string  patterns
 * @param  string  value
 * @return bool
 */
func (this *str) Is(patterns []string, value ...string) bool {
	value = append(value, this._str)
	if len(patterns) == 0 {
		return false
	}
	for _, pattern := range patterns {
		if pattern == value[0] {
			return true
		}
		return regexp.MustCompile(fmt.Sprintf(`^%s\z`, strings.ReplaceAll(regexp.QuoteMeta(pattern), `\*`, `.*`))).MatchString(value[0])
	}
	return false
}

/**
 * Return the length of the given string.
 *
 * @param  string  value
 * @return int
 */
func (this *str) Length(value ...string) int {
	value = append(value, this._str)
	return utf8.RuneCountInString(value[0])
}
