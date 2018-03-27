package controllers

import (
	"github.com/KeKsBoTer/socialloot/models"
)

type UserController struct {
	AuthController
}

func (this *UserController) Get() {
	userName := this.Ctx.Input.Param(":user")
	user := &models.User{
		Name: userName,
	}
	if err := user.Read("Name"); err != nil {
		this.Abort("404")
		return
	}
	this.Data["User"] = user
	this.TplName = "pages/users/page.tpl"
}
