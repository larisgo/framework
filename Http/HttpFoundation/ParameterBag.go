package HttpFoundation

type ParameterBag struct {
	parameters map[string][]string
}

func NewParameterBag(parameters map[string][]string) *ParameterBag {
	return &ParameterBag{parameters: parameters}
}

/**
 * Returns the parameters.
 *
 * @return array An array of parameters
 */
func (this *ParameterBag) All() map[string][]string {
	return this.parameters
}

/**
 * Returns the parameter keys.
 *
 * @return array An array of parameter keys
 */
func (this *ParameterBag) Keys() []string {
	keys := []string{}
	for k, _ := range this.parameters {
		keys = append(keys, k)
	}
	return keys
}

/**
 * Replaces the current parameters by a new set.
 *
 * @param array parameters An array of parameters
 */
func (this *ParameterBag) Replace(parameters map[string][]string) {
	this.parameters = parameters
}

/**
 * Adds parameters.
 *
 * @param array parameters An array of parameters
 */
func (this *ParameterBag) Add(parameters map[string][]string) {
	for k, v := range parameters {
		this.parameters[k] = v
	}
}

/**
 * Returns a parameter by name.
 *
 * @param string key     The key
 * @param mixed  default The default value if the parameter key does not exist
 *
 * @return mixed
 */
func (this *ParameterBag) Get(key string, _default string) string {
	value := this.Gets(key)
	l := len(value)
	if l == 0 {
		return _default
	}
	return value[l-1]
}

/**
 * Returns a parameter by name.
 *
 * @param string key     The key
 * @param mixed  default The default value if the parameter key does not exist
 *
 * @return mixed
 */
func (this *ParameterBag) Gets(key string, _default ...[]string) []string {
	_default = append(_default, []string{})
	if v, ok := this.parameters[key]; ok {
		return v
	}
	return _default[0]
}

/**
 * Sets a parameter by name.
 *
 * @param string key   The key
 * @param mixed  value The value
 */
func (this *ParameterBag) Set(key string, value []string) {
	this.parameters[key] = value
}

/**
 * Returns true if the parameter is defined.
 *
 * @param string $key The key
 *
 * @return bool true if the parameter exists, false otherwise
 */
func (this *ParameterBag) Has(key string) bool {
	_, ok := this.parameters[key]
	return ok
}

/**
 * Removes a parameter.
 *
 * @param string $key The key
 */
func (this *ParameterBag) Remove(key string) {
	delete(this.parameters, key)
}
