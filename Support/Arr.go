package Support

import (
	"strings"
)

type arr struct {
}

var Arr func() *arr = func() *arr {
	return &arr{}
}

/**
 * Determine whether the given value is array accessible.
 *
 * @param  mixed  value
 * @return bool
 */
func (this *arr) Accessible(value interface{}) bool {
	_, ok := value.(map[string]interface{})
	return ok
}

/**
 * Add an element to an array using "dot" notation if it doesn't exist.
 *
 * @param  array   array
 * @param  string  key
 * @param  mixed   value
 * @return array
 */
func (this *arr) Add(array map[string]interface{}, key string, value interface{}) map[string]interface{} {
	if this.Get(array, key) == nil {
		this.Set(&array, key, value)
	}

	return array
}

/**
 * Determine if the given key Exists in the provided array.
 *
 * @param  \ArrayAccess|array  array
 * @param  string|int  key
 * @return bool
 */
func (this *arr) Exists(array map[string]interface{}, key string) bool {
	_, ok := array[key]
	return ok
}

/**
 * Get an item from an array using "dot" notation.
 *
 * @param  \ArrayAccess|array  array
 * @param  string  key
 * @param  mixed   default
 * @return mixed
 */
func (this *arr) Get(array map[string]interface{}, key string, _default ...interface{}) interface{} {
	_default = append(_default, nil)

	if this.Exists(array, key) {
		return array[key]
	}

	if strings.Index(key, ".") == -1 {
		if v, ok := array[key]; ok {
			return v
		} else {
			return _default[0]
		}
	}

	var _array interface{} = array
	for _, segment := range strings.Split(key, ".") {
		if this.Accessible(_array) && this.Exists(_array.(map[string]interface{}), segment) {
			_array = _array.(map[string]interface{})[segment]
		} else {
			return _default[0]
		}
	}
	return _array
}

/**
 * Check if an item or items exist in an array using "dot" notation.
 *
 * @param  \ArrayAccess|array  array
 * @param  string|array  keys
 * @return bool
 */
func (this *arr) Has(array map[string]interface{}, keys []string) bool {

	if len(keys) == 0 {
		return false
	}
	for _, key := range keys {
		var subKeyArray interface{} = array
		if this.Exists(array, key) {
			continue
		}
		for _, segment := range strings.Split(key, ".") {
			if this.Accessible(subKeyArray) && this.Exists(subKeyArray.(map[string]interface{}), segment) {
				subKeyArray = subKeyArray.(map[string]interface{})[segment]
			} else {
				return false
			}
		}
	}

	return true
}

/**
 * Set an array item to a given value using "dot" notation.
 *
 * If no key is given to the method, the entire array will be replaced.
 *
 * @param  array   array
 * @param  string  key
 * @param  mixed   value
 * @return array
 */
func (this *arr) Set(array *map[string]interface{}, key string, value interface{}) *map[string]interface{} {

	keys := strings.Split(key, ".")
	for len(keys) > 1 {
		key := keys[0]
		keys = keys[1:]

		// If the key doesn't exist at this depth, we will just create an empty array
		// to hold the next value, allowing us to create the arrays to hold final
		// values at the correct depth. Then we'll keep digging into the array.
		if v, ok := (*array)[key]; !ok {
			(*array)[key] = map[string]interface{}{}
		} else if _, tok := v.(map[string]interface{}); !tok {
			(*array)[key] = map[string]interface{}{}
		}
		(*array) = (*array)[key].(map[string]interface{})
	}

	(*array)[keys[0]] = value
	return array
}
