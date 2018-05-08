package controllers

import (
	"net/http"

	"github.com/KeKsBoTer/socialloot/lib"
	"github.com/KeKsBoTer/socialloot/models"
)

// LoginController provides login and signup page
type LoginController struct {
	AuthController
}

// LoginPage serves login page and redirects to root page if the user is allready logged in
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

// Login authenticates user and sets session key
func (c *LoginController) Login() {
	form := &models.LoginForm{}
	handleForm(form, &c.AuthController, func(r *APIResponse) {
		user, err := lib.Authenticate(form.UserName, form.Password)
		if err != nil {
			r.Fail("", err)
			return
		}
		c.SetLogin(user)
		r.Success = true
	})
}

// Logout deletes session
func (c *LoginController) Logout() {
	c.DelLogin()
	c.RedirectForm()
}

// SignupPage serves signup page
func (c *LoginController) SignupPage() {
	c.TplName = "pages/login/signup.tpl"
	c.Data["Title"] = "Sign up to Socialloot"
}

// Signup handles register requests
func (c *LoginController) Signup() {
	form := &models.SignUpForm{}
	handleForm(form, &c.AuthController, func(r *APIResponse) {
		user, err := lib.SignupUser(form.UserName, form.Password)
		if err != nil {
			r.Fail("", err)
			return
		}
		c.SetLogin(user)
		r.Success = true
	})
}
