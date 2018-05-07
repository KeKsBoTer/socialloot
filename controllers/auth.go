package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/KeKsBoTer/socialloot/models"
	"github.com/astaxie/beego"
)

type NestPreparer interface {
	NestPrepare()
}

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

		// redirect destination in http get request
		if dst := c.GetString("dest"); len(dst) > 1 {
			c.Data["Dest"] = dst
		}

		c.Data["xsrf_token"] = c.XSRFToken()

		c.Data["URL"] = c.Ctx.Input.URI()
		var topics []*models.Topic
		if _, err := models.Topics().OrderBy("name").Limit(20).All(&topics); err != nil {
			c.Abort("500")
			return
		}
		c.Data["Topics"] = topics

		c.Layout = "base.tpl"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["BaseHeader"] = "components/header.tpl"
	}
	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
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

// Prepare checks if the user is authorized.
// If not an error is returned
// The user is redirected to the login page
func (c *NeedsAuthController) Prepare() {
	isLogin := c.GetSession(UserInfoKey) != nil
	if !isLogin {
		switch {
		case c.Ctx.Input.IsGet():
			c.RedirectForm()
			if !c.Ctx.Output.IsRedirect() {
				c.Ctx.Redirect(http.StatusSeeOther, c.LoginPath())
			}
		case c.Ctx.Input.IsPost():
			r := APIResponse{
				Success: false,
				Message: "unauthorized",
				Dest:    c.URLFor("LoginController.LoginPage"),
			}
			j, _ := json.Marshal(r)
			c.CustomAbort(http.StatusUnauthorized, string(j))
		default:
			// only HTTP get and post are allowed
			c.Abort("405")
		}
	} else {
		c.AuthController.Prepare()
	}
}
