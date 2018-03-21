package routers

import (
	ctl "github.com/KeKsBoTer/socialloot/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &ctl.UsersController{}, "get:Index")
	beego.Router("/login", &ctl.LoginController{}, "get,post:Login")
	beego.Router("/logout", &ctl.LoginController{}, "get:Logout")
	beego.Router("/signup", &ctl.LoginController{}, "get,post:Signup")
}
