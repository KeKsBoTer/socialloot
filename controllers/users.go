package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

type UsersController struct {
	BaseController
}

func (c *UsersController) NestPrepare() {
	if !c.IsLogin {
		c.Ctx.Redirect(http.StatusFound, c.LoginPath())
		return
	}
}

// func (c *UsersController) NestFinish() {}

func (c *UsersController) Index() {
	beego.ReadFromRequest(&c.Controller)

	c.TplName = "users/index.tpl"
}
