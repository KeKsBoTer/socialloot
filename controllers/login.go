package controllers

import (
	"net/http"

	"github.com/KeKsBoTer/socialloot/lib"
	"github.com/KeKsBoTer/socialloot/models"
)

type LoginController struct {
	AuthController
}

func (c *LoginController) LoginPage() {
	if c.IsLogin() {
		c.RedirectForm()
		if !c.Ctx.Output.IsRedirect() {
			c.Ctx.Redirect(http.StatusSeeOther, c.LoginPath())
		}
		return
	}
	c.TplName = "pages/login/login.tpl"
	c.Data["Title"] = "Login"
}

func (c *LoginController) Login() {
	form := &models.LoginForm{}
	handleForm(form, &c.AuthController, func(r *ApiResponse) {
		user, err := lib.Authenticate(form.UserName, form.Password)
		if err != nil {
			r.Fail("", err)
			return
		}
		c.SetLogin(user)
		r.Success = true
	})
}

func (c *LoginController) Logout() {
	c.DelLogin()
	c.RedirectForm()
}

func (c *LoginController) SignupPage() {
	c.TplName = "pages/login/signup.tpl"
	c.Data["Title"] = "Sign up to Socialloot"
}

func (c *LoginController) Signup() {
	form := &models.SignUpForm{}
	handleForm(form, &c.AuthController, func(r *ApiResponse) {
		user, err := lib.SignupUser(form.UserName, form.Password)
		if err != nil {
			r.Fail("", err)
			return
		}
		c.SetLogin(user)
		r.Success = true
	})
}
