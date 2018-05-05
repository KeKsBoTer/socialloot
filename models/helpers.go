package models

import (
	"github.com/astaxie/beego/validation"
)

// IsValid checks if a given model passes the validation and returns the first occurring error
func IsValid(model interface{}) error {
	valid := validation.Validation{}
	b, err := valid.Valid(model)
	if !b {
		return valid.Errors[0]
	}
	return err
}
