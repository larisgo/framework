package Config

import (
	"github.com/larisgo/framework/Support"
)

type Repository struct {
	items map[string]interface{}
}

func NewRepository() *Repository {
	return &Repository{items: map[string]interface{}{}}
}

/**
 * Determine if the given configuration value exists.
 *
 * @param  string  key
 * @return bool
 */
func (this *Repository) Has(key string) bool {
	return Support.Arr().Has(this.items, []string{key})
}

/**
 * Get the specified configuration value.
 *
 * @param  array|string  key
 * @param  mixed   default
 * @return mixed
 */
func (this *Repository) Get(key string, _default ...interface{}) interface{} {
	return Support.Arr().Get(this.items, key, _default...)
}

/**
 * Get many configuration values.
 *
 * @param  array  keys
 * @return array
 */
func (this *Repository) GetMany(keys map[string]interface{}) map[string]interface{} {
	config := map[string]interface{}{}
	for key, _default := range keys {
		config[key] = Support.Arr().Get(this.items, key, _default)
	}

	return config
}

/**
 * Set a given configuration value.
 *
 * @param  array|string  key
 * @param  mixed   value
 * @return void
 */
func (this *Repository) Set(key string, value ...interface{}) {
	value = append(value, nil)
	Support.Arr().Set(&this.items, key, value[0])
}

/**
 * Get all of the configuration items for the application.
 *
 * @return array
 */
func (this *Repository) All() map[string]interface{} {
	return this.items
}
