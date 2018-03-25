package models

import (
	"errors"

	"github.com/astaxie/beego"
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
			beego.Warning(err.Key, ":", err.Message)
			msg += err.Key + " : " + err.Message
		}
		return errors.New(msg)
	}
	return err
}
