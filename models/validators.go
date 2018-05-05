package models

import (
	"net/url"

	"github.com/astaxie/beego/validation"
)

func (form *SignUpForm) Valid(v *validation.Validation) {
	if form.Password != form.PasswordRe {
		v.AddError("PasswordRe.Match", "Password and reentered password do not match")
	}
}

func (form *SubmitForm) Valid(v *validation.Validation) {
	switch form.Type {
	case PostTypeLink:
		v.Required(form.Content, "Content.Required").Message("Please enter a link")
		v.MaxSize(form.Content, 40000, "Content.MaxSize")
		if u, err := url.ParseRequestURI(form.Content); err != nil || !u.IsAbs() {
			v.AddError("Content.Match", "Invalid url format for link")
		}
	case PostTypeText:
		v.Required(form.Content, "Content.Required").Message("Please enter a text")
		v.MaxSize(form.Content, 40000, "Content.MaxSize")
	case PostTypeImage:
	default:
		v.AddError("Type", "Invalid post type")

	}
}

func (form *VoteForm) Valid(v *validation.Validation) {
	if form.Direction != VoteDirectionUp && form.Direction != VoteDirectionDown {
		v.SetError("Direction.Match", "Direction must be up or downvote")
	}
}
