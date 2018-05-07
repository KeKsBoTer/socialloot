package models

import (
	"net/url"

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

// Valid checks if the entered password and the reentered passwords are the same
func (form *SignUpForm) Valid(v *validation.Validation) {
	if form.Password != form.PasswordRe {
		v.AddError("PasswordRe.Match", "Password and reentered password do not match")
	}
}

// Valid checks the content based on the post type:
// 	- text: is the text's length smaller than 40000 characters?
//  - link: is the enterd link a valid url?
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

// Valid checks if the vote direction is up or down
func (form *VoteForm) Valid(v *validation.Validation) {
	if form.Direction != VoteDirectionUp && form.Direction != VoteDirectionDown {
		v.SetError("Direction.Match", "Direction must be up or downvote")
	}
}
