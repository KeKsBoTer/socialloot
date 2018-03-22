package routers

import (
	"strings"

	ctl "github.com/KeKsBoTer/socialloot/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// AuthFilter redirects to the login page if the user is not authenticated
func AuthFilter(ctx *context.Context) {
	if strings.HasPrefix(ctx.Input.URL(), "/login") {
		return
	}

	_, ok := ctx.Input.Session("userinfo").(int)
	if !ok {
		ctx.Redirect(302, "/login")
	}
}

func init() {

	beego.InsertFilter("*", beego.BeforeExec, AuthFilter)

	beego.Router("/", &ctl.UsersController{}, "get:Index")
	beego.Router("/login", &ctl.LoginController{}, "get,post:Login")
	beego.Router("/logout", &ctl.LoginController{}, "get:Logout")
	beego.Router("/signup", &ctl.LoginController{}, "get,post:Signup")
	beego.Router("/submit", &ctl.SubmitController{}, "get,post:Submit")
	beego.Router("/createtopic", &ctl.SubmitController{}, "get,post:CreateTopic")

	beego.Router("/topic/:topic:string", &ctl.TopicController{}, "get:PostsList")
}
