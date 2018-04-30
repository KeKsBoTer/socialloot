package models

import (
	"github.com/astaxie/beego/validation"
)

func (form *SignUpForm) Valid(v *validation.Validation) {
	if form.Password != form.PasswordRe {
		v.AddError("Repassword", "Does not matched password, repassword")
	}
}

func (form *SubmitForm) Valid(v *validation.Validation) {
	if form.Type == PostTypeLink || form.Type == PostTypeText {
		v.MaxSize(form, 40000, form.Content)
	} else if form.Type != PostTypeImage {
		v.AddError("Type", "Invalid post type")
	}
}

func (form *VoteForm) Valid(v *validation.Validation) {
	if form.Direction != VoteDirectionUp && form.Direction != VoteDirectionDown {
		v.SetError("direction", "Direction must be up or downvote")
	}
}
