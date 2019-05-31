package Routing

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

/**
 * This string defines the characters that are automatically considered separators in front of
 * optional placeholders (with default and no static text following). Such a single separator
 * can be left out together with the optional placeholder from matching and generating URLs.
 */
// const REGEX_DELIMITER = `#`

/**
 * The maximum supported length of a PCRE subpattern name
 * http://pcre.org/current/doc/html/pcre2pattern.html#SEC16.
 *
 * @internal
 */
const SEPARATORS = `/,;.:-_~+*=@|`

/**
 * {@inheritdoc}
 *
 * @throws \InvalidArgumentException if a path variable is named _fragment
 * @throws \LogicException           if a variable is referenced more than once
 * @throws \DomainException          if a variable name starts with a digit or if it is too long to be successfully used as
 *                                   a PCRE subpattern
 */
const VARIABLE_MAXIMUM_LENGTH = 32

const INT_MAX = int(^uint(0) >> 1)

type SymfonyRouteCompiler struct {
}

type compilePatternResult struct {
	staticPrefix string
	regex        string
	tokens       [][]string
	variables    map[string]string
}

func NewSymfonyRouteCompiler() *SymfonyRouteCompiler {
	return &SymfonyRouteCompiler{}
}

func (this *SymfonyRouteCompiler) Compile(route *SymfonyRoute) *CompiledRoute {
	hostVariables := map[string]string{}
	variables := map[string]string{}
	hostRegex := ""
	hostTokens := [][]string{}

	if host := route.GetHost(); host != "" {
		result := this.compilePattern(route, host, true)

		hostVariables = result.variables
		variables = hostVariables

		hostTokens = result.tokens
		hostRegex = result.regex
	}

	path := route.GetPath()

	result := this.compilePattern(route, path, false)

	staticPrefix := result.staticPrefix

	pathVariables := result.variables

	for _, pathParam := range pathVariables {
		if pathParam == "_fragment" {
			panic(errors.New(fmt.Sprintf(`Route pattern "%s" cannot contain "_fragment" as a path parameter.`, route.GetPath())))
		}
	}

	for k, v := range pathVariables {
		variables[k] = v
	}

	tokens := result.tokens
	_regex := result.regex

	return NewCompiledRoute(
		staticPrefix,
		_regex,
		tokens,
		pathVariables,
		hostRegex,
		hostTokens,
		hostVariables,
		variables, // map已经过滤
	)
}

func (this *SymfonyRouteCompiler) compilePattern(route *SymfonyRoute, pattern string, isHost bool) *compilePatternResult {
	tokens := [][]string{}
	variables := map[string]string{}
	// matches := []string{}
	pos := 0
	var defaultSeparator string
	if isHost {
		defaultSeparator = "."
	} else {
		defaultSeparator = "/"
	}
	useUtf8 := utf8.ValidString(pattern)
	needsUtf8 := route.GetOption("utf8").(bool)
	if !needsUtf8 && useUtf8 && regexp.MustCompile(`[\x80-\xFF]`).MatchString(pattern) {
		panic(errors.New(fmt.Sprintf(`Cannot use UTF-8 route patterns without setting the "utf8" option for route "%s".`, route.GetPath())))
	}
	if !useUtf8 && needsUtf8 {
		panic(errors.New(fmt.Sprintf(`Cannot mix UTF-8 requirements with non-UTF-8 pattern "%s".`, pattern)))
	}
	// Match all variables enclosed in "{}" and iterate over them. But we only want to match the innermost variable
	// in case of nested "{}", e.g. {foo{bar}}. This in ensured because \w does not match "{" or "}" itself.
	for _, matche := range regexp.MustCompile(`\{\w+\}`).FindAllStringIndex(pattern, -1) {
		varName := pattern[matche[0]+1 : matche[1]-1]
		// get all static text preceding the current variable
		precedingText := pattern[pos:matche[0]]
		pos = matche[1]

		var precedingChar string
		if len(precedingText) == 0 {
			precedingChar = ""
		} else if useUtf8 {
			_precedingChar := []rune(precedingText)
			precedingChar = string(_precedingChar[len(_precedingChar)-1:])
		} else {
			precedingChar = precedingText[len(precedingText)-1:]
		}

		var isSeparator bool
		if precedingChar != "" && strings.Index(SEPARATORS, precedingChar) > -1 {
			isSeparator = true
		} else {
			isSeparator = false
		}

		// A PCRE subpattern name must start with a non-digit. Also a PHP variable cannot start with a digit so the
		// variable would not be usable as a Controller action argument.
		if regexp.MustCompile(`^\d`).MatchString(varName) {
			panic(errors.New(fmt.Sprintf(`Variable name "%s" cannot start with a digit in route pattern "%s". Please use a different name.`, varName, pattern)))
		}
		if _, ok := variables[varName]; ok {
			panic(errors.New(fmt.Sprintf(`Route pattern "%s" cannot reference variable name "%s" more than once.`, pattern, varName)))
		}

		if len(varName) > VARIABLE_MAXIMUM_LENGTH {
			panic(errors.New(fmt.Sprintf(`Variable name "%s" cannot be longer than %s characters in route pattern "%s". Please use a shorter name.`, pattern)))
		}

		if isSeparator && precedingText != precedingChar {
			tokens = append(tokens, []string{"text", precedingText[0 : len(precedingText)-len(precedingChar)]})
		} else if !isSeparator && len(precedingText) > 0 {
			tokens = append(tokens, []string{"text", precedingText})
		}

		_regexp := route.GetRequirement(varName)
		if _regexp == "" {
			followingPattern := pattern[pos:]
			// Find the next static character after the variable that functions as a separator. By default, this separator and '/'
			// are disallowed for the variable. This default requirement makes sure that optional variables can be matched at all
			// and that the generating-matching-combination of URLs unambiguous, i.e. the params used for generating the URL are
			// the same that will be matched. Example: new Route('/{page}.{_format}', ['_format' => 'html'])
			// If {page} would also match the separating dot, {_format} would never match as {page} will eagerly consume everything.
			// Also even if {_format} was not optional the requirement prevents that {page} matches something that was originally
			// part of {_format} when generating the URL, e.g. _format = 'mobile.html'.
			nextSeparator := this.findNextSeparator(followingPattern, useUtf8)
			_regexp = fmt.Sprintf(
				`[^%s%s]+`,
				regexp.QuoteMeta(defaultSeparator),
				func() string {
					if defaultSeparator != nextSeparator && nextSeparator != "" {
						return regexp.QuoteMeta(nextSeparator)
					}
					return ""
				}(),
			)
			// golang not support
			/**
			if (nextSeparator != "" && !regexp.MustCompile(`^\{\w+\}`).MatchString(followingPattern)) || followingPattern == "" {
				// When we have a separator, which is disallowed for the variable, we can optimize the regex with a possessive
				// quantifier. This prevents useless backtracking of PCRE and improves performance by 20% for matching those patterns.
				// Given the above example, there is no point in backtracking into {page} (that forbids the dot) when a dot must follow
				// after it. This optimization cannot be applied when the next char is no real separator or when the next variable is
				// directly adjacent, e.g. '/{x}{y}'.
				_regexp += "+"
			}
			*/
		} else {
			if !utf8.ValidString(_regexp) {
				useUtf8 = false
			} else if !needsUtf8 && regexp.MustCompile(`[\x80-\xFF]|(?<!\\)\\(?:\\\\)*+(?-i:X|[pP][\{CLMNPSZ]|x\{[A-Fa-f0-9]{3})`).MatchString(_regexp) {
				panic(errors.New(fmt.Sprintf(`Cannot use UTF-8 route requirements without setting the "utf8" option for variable "%s" in pattern "%s".`, varName, pattern)))
			}
			if !useUtf8 && needsUtf8 {
				panic(errors.New(fmt.Sprintf(`Cannot mix UTF-8 requirement with non-UTF-8 charset for variable "%s" in pattern "%s"`, varName, pattern)))
			}
			_regexp = this.transformCapturingGroupsToNonCapturings(_regexp)
		}
		tokens = append(tokens, []string{"variable", func() string {
			if isSeparator {
				return precedingChar
			}
			return ""
		}(), _regexp, varName})
		variables[varName] = ""
	}

	if pos < len(pattern) {
		tokens = append(tokens, []string{"text", pattern[pos:]})
	}
	// find the first optional token
	firstOptional := INT_MAX
	if !isHost {
		for i := len(tokens) - 1; i >= 0; i -= 1 {
			if token := tokens[i]; token[0] == "variable" && route.HasDefault(token[3]) {
				firstOptional = i
			} else {
				break
			}
		}
	}

	// compute the matching regexp
	_regexp := ""
	nbToken := len(tokens)
	for i := 0; i < nbToken; i += 1 {
		_regexp += this.computeRegexp(tokens, i, firstOptional)
	}
	// golang not support D `(?sD%s)^%s$`
	// D (PCRE_DOLLAR_ENDONLY)
	// If this modifier is set, a dollar metacharacter in the pattern matches only at the end of the subject string. Without this modifier, a dollar also matches immediately before the final character if it is a newline (but not before any other newlines). This modifier is ignored if m modifier is set. There is no equivalent to this modifier in Perl.
	_regexp = fmt.Sprintf(`(?s%s)^%s$`, func() string {
		if isHost {
			return "i"
		}
		return ""
	}(), _regexp)

	// enable Utf8 matching if really required
	if needsUtf8 {
		// _regexp += "u"; // golang not support
		nbToken := len(tokens)
		for i := 0; i < nbToken; i += 1 {
			if tokens[i][0] == "variable" {
				tokens[i] = append(tokens[i], "true")
			}
		}
	}
	return &compilePatternResult{
		staticPrefix: this.determineStaticPrefix(route, tokens),
		regex:        _regexp,
		tokens:       tokens,
		variables:    variables,
	}
}

/**
 * Determines the longest static prefix possible for a route.
 */
func (this *SymfonyRouteCompiler) determineStaticPrefix(route *SymfonyRoute, tokens [][]string) string {
	if tokens[0][0] != "text" {
		if route.HasDefault(tokens[0][3]) || "/" == tokens[0][1] {
			return ""
		}
		return tokens[0][1]
	}

	prefix := tokens[0][1]
	if len(tokens) > 1 && len(tokens[1]) > 1 && tokens[1][1] != "/" && route.HasDefault(tokens[1][3]) {
		prefix += tokens[1][1]
	}

	return prefix
}

/**
 * Returns the next static character in the Route pattern that will serve as a separator (or the empty string when none available).
 */
func (this *SymfonyRouteCompiler) findNextSeparator(pattern string, useUtf8 bool) string {
	if pattern == "" {
		// return empty string if pattern is empty or false (false which can be returned by substr)
		return ""
	}
	// first remove all placeholders from the pattern so we can find the next real static character
	if pattern = regexp.MustCompile(`\{\w+\}`).ReplaceAllString(pattern, ``); pattern == "" {
		return ""
	}
	if useUtf8 {
		pattern = string([]rune(pattern)[0:1])
	}
	if strings.Index(SEPARATORS, pattern) > -1 {
		return pattern
	}
	return ""
}

/**
 * Computes the regexp used to match a specific token. It can be static text or a subpattern.
 *
 * @param array tokens        The route tokens
 * @param int   index         The index of the current token
 * @param int   firstOptional The index of the first optional token
 *
 * @return string The regexp pattern for a single token
 */
func (this *SymfonyRouteCompiler) computeRegexp(tokens [][]string, index int, firstOptional int) string {
	token := tokens[index]
	if token[0] == "text" {
		// Text tokens
		return regexp.QuoteMeta(token[1])
	} else {
		// Variable tokens
		if index == 0 && firstOptional == 0 {
			// When the only token is an optional variable token, the separator is required
			return fmt.Sprintf(`%s(?P<%s>%s)?`, regexp.QuoteMeta(token[1]), token[3], token[2])
		} else {
			_regexp := fmt.Sprintf(`%s(?P<%s>%s)`, regexp.QuoteMeta(token[1]), token[3], token[2])
			if index >= firstOptional {
				// Enclose each optional token in a subpattern to make it optional.
				// "?:" means it is non-capturing, i.e. the portion of the subject string that
				// matched the optional subpattern is not passed back.
				_regexp = "(?:" + _regexp
				nbTokens := len(tokens)
				if nbTokens-1 == index {
					// Close the optional subpatterns
					_regexp += strings.Repeat(`)?`, nbTokens-firstOptional-(func() int {
						if firstOptional == 0 {
							return 1
						}
						return 0
					}()))
				}
			}

			return _regexp
		}
	}
}

func (this *SymfonyRouteCompiler) transformCapturingGroupsToNonCapturings(regex string) string {
	_regexp := []rune(regex)
	for i := 0; i < len(_regexp); i += 1 {
		if string(_regexp[i]) == `\` {
			i += 1
			continue
		}
		if string(_regexp[i]) != `(` || (i+2 > len(_regexp)) {
			continue
		}
		if i += 1; string(_regexp[i]) == `*` || string(_regexp[i]) == `?` {
			i += 1
			continue
		}
		regex = regex[:i] + "?:" + regex[i:]
		i += 1
	}

	return regex
}
