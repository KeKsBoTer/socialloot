package models

import (
	"errors"

	"github.com/astaxie/beego/validation"
)

/*
	Get form error
*/
func IsValid(model interface{}) (err error) {
	valid := validation.Validation{}
	b, err := valid.Valid(model)
	if !b {
		var msg string
		for _, err := range valid.Errors {
			if len(err.Field) > 1 {
				msg += err.Field + ": "
			}
			msg += err.Message
			// only diplay first error to user
			break
		}
		return errors.New(msg)
	}
	return err
}
