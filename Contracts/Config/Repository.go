package Config

type Repository interface {
	Has(string) bool
	Get(string, ...interface{}) interface{}
	GetMany(map[string]interface{}) map[string]interface{}
	Set(string, ...interface{})
	All() map[string]interface{}
}
