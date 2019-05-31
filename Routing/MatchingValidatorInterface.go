package Routing

type ValidatorInterface interface {
	matches(*Route, string) bool
}
