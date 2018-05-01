package models

import (
	"github.com/astaxie/beego/validation"
)

/*
	Get form error
*/
func IsValid(model interface{}) (err error) {
	valid := validation.Validation{}
	b, err := valid.Valid(model)
	if !b {
		return valid.Errors[0]
	}
	return err
}
