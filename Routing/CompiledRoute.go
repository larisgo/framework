package Routing

type CompiledRoute struct {
	variables     map[string]string
	tokens        [][]string
	staticPrefix  string
	regex         string
	pathVariables map[string]string
	hostVariables map[string]string
	hostRegex     string
	hostTokens    [][]string
}

func NewCompiledRoute(staticPrefix string, regex string, tokens [][]string, pathVariables map[string]string, hostRegex string, hostTokens [][]string, hostVariables map[string]string, variables map[string]string) (this *CompiledRoute) {
	this = &CompiledRoute{}
	this.staticPrefix = staticPrefix
	this.regex = regex
	this.tokens = tokens
	this.pathVariables = pathVariables
	this.hostRegex = hostRegex
	this.hostTokens = hostTokens
	this.hostVariables = hostVariables
	this.variables = variables

	return this
}

/**
 * Returns the static prefix.
 *
 * @return string The static prefix
 */
func (this *CompiledRoute) GetStaticPrefix() string {
	return this.staticPrefix
}

/**
 * Returns the regex.
 *
 * @return string The regex
 */
func (this *CompiledRoute) GetRegex() string {
	return this.regex
}

/**
 * Returns the host regex.
 *
 * @return string|null The host regex or null
 */
func (this *CompiledRoute) GetHostRegex() string {
	return this.hostRegex
}

/**
 * Returns the tokens.
 *
 * @return array The tokens
 */
func (this *CompiledRoute) GetTokens() [][]string {
	return this.tokens
}

/**
 * Returns the host tokens.
 *
 * @return array The tokens
 */
func (this *CompiledRoute) GetHostTokens() [][]string {
	return this.hostTokens
}

/**
 * Returns the variables.
 *
 * @return array The variables
 */
func (this *CompiledRoute) GetVariables() map[string]string {
	return this.variables
}

/**
 * Returns the path variables.
 *
 * @return array The variables
 */
func (this *CompiledRoute) GetPathVariables() map[string]string {
	return this.pathVariables
}

/**
 * Returns the host variables.
 *
 * @return array The variables
 */
func (this *CompiledRoute) GetHostVariables() map[string]string {
	return this.hostVariables
}
