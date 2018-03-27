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
	// server answer as json
	r := apiResponse(&c.Controller)
	defer c.ServeJSON(true)

	name := c.GetString("Name")
	password := c.GetString("Password")

	user, err := lib.Authenticate(name, password)
	if err != nil {
		r.Success = false
		r.Message = err.Error()
		return
	}
	c.SetLogin(user)
	r.Success = true
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
	// server answer as json
	r := apiResponse(&c.Controller)
	defer c.ServeJSON(true)

	u := &models.User{}
	if err := c.ParseForm(u); err != nil {
		r.Success = false
		r.Message = "Signup invalid!"
		return
	}
	if err := models.IsValid(u); err != nil {
		r.Success = false
		r.Message = err.Error()
		return
	}

	if err := lib.SignupUser(u); err != nil {
		r.Success = false
		r.Message = err.Error()
		return
	}
	c.SetLogin(u)
	r.Success = true
}
