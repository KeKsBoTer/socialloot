package controllers

import (
	"net/http"

	"github.com/KeKsBoTer/socialloot/models"
	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller

	User *models.User
}

// UserInfoKey is the session key for the user id
const UserInfoKey = "userinfo"

func (c *AuthController) Prepare() {
	isLogin := c.GetSession(UserInfoKey) != nil
	if isLogin {
		c.User = c.GetLogin()
	}

	// set data for html rendering
	if c.Ctx.Input.IsGet() {
		c.Data["IsLogin"] = isLogin
		c.Data["User"] = c.User

		c.Data["HeadStyles"] = []string{}
		c.Data["HeadScripts"] = []string{}

		c.Layout = "base.tpl"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["BaseHeader"] = "header.tpl"
	}
}

func (c *AuthController) RedirectForm() {
	if dst := c.GetString("dest"); len(dst) > 0 {
		c.Redirect(dst, http.StatusSeeOther)
	}
}

func (c *AuthController) GetLogin() *models.User {
	if i, ok := c.GetSession(UserInfoKey).(int); ok {
		u := &models.User{
			Id: i,
		}
		u.Read()
		return u
	}
	return nil
}

func (c *AuthController) IsLogin() bool {
	return c.User != nil
}

func (c *AuthController) DelLogin() {
	c.DelSession(UserInfoKey)
}

func (c *AuthController) SetLogin(user *models.User) {
	c.SetSession(UserInfoKey, user.Id)
}

func (c *AuthController) LoginPath() string {
	return c.URLFor("LoginController.Login")
}

// NeedsAuthController redirects to Login page if user is not authenticated
type NeedsAuthController struct {
	AuthController
}

func (c *NeedsAuthController) Prepare() {
	c.AuthController.Prepare()
	if !c.IsLogin() {
		if c.Ctx.Input.IsGet() {
			c.RedirectForm()
			if !c.Ctx.Output.IsRedirect() {
				c.Ctx.Redirect(http.StatusSeeOther, c.URLFor("IndexController.Index"))
			}
		} else {
			c.Abort("401")
		}
	}
}
