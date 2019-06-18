package Support

import (
	"fmt"
	"regexp"
	"strings"
)

type str struct {
	snakeCache  []string
	camelCache  []string
	studlyCache []string
}

var Str func() *str = func() *str {
	return &str{}
}

/**
 * Determine if a given string matches a given pattern.
 *
 * @param  []string  patterns
 * @param  string  value
 * @return bool
 */
func (this *str) Is(value string, patterns []string) bool {
	if len(patterns) == 0 {
		return false
	}
	for _, pattern := range patterns {
		if pattern == value {
			return true
		}
		return regexp.MustCompile(fmt.Sprintf(`^%s\z`, strings.ReplaceAll(regexp.QuoteMeta(pattern), `\*`, `.*`))).MatchString(value)
	}
	return false
}

/**
 * Return the length of the given string.
 *
 * @param  string  value
 * @return int
 */
func (this *str) Length(value string) int {
	return len([]rune(value))
}

/**
 * Determine if a given string contains a given substring.
 *
 * @param  []string  needles
 * @param  string  haystack
 * @return bool
 */
func (this *str) Contains(haystack string, needles []string) bool {

	for _, needle := range needles {
		if needle != "" && strings.Index(haystack, needle) > 0 {
			return true
		}
	}

	return false
}

/**
 * Returns the portion of string specified by the start and length parameters.
 *
 * @param  string  string
 * @param  int  start
 * @param  int|null  length
 * @return string
 */
func (this *str) Substr(str string, start int, length ...int) string {
	length = append(length, -1)
	_str := []rune(str)
	_str_length := len(_str)
	end := start + length[0]
	if (start > _str_length) || (start < 0) {
		start = 0
	}
	if (end < start) || (end > _str_length) {
		end = _str_length
	}

	return string(str[start:end])
}
