package models

import (
	"github.com/astaxie/beego/validation"
)

func (form *SignUpForm) Valid(v *validation.Validation) {
	if form.Password != form.PasswordRe {
		v.AddError("Repassword", "Does not matched password, repassword")
	}
}
