package server

import "gopkg.in/go-playground/validator.v9"

// https://github.com/go-playground/validator
var validate *validator.Validate

// 单例
func ValidatorInstance() *validator.Validate {
	if validate == nil {
		validate = validator.New()
	}
	return validate
}
