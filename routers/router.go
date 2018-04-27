package routers

import (
	ctl "github.com/KeKsBoTer/socialloot/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.SetStaticPath("/favicon.ico", "/static/img/link_icon.png") // relative path

	beego.Router("/?:choice", &ctl.IndexController{})
	beego.Router("/login", &ctl.LoginController{}, "get:LoginPage;post:Login")
	beego.Router("/logout", &ctl.LoginController{}, "get:Logout")
	beego.Router("/signup", &ctl.LoginController{}, "get:SignupPage;post:Signup")
	beego.Router("/submit", &ctl.SubmitController{}, "get,post:Submit")
	beego.Router("/createtopic", &ctl.SubmitController{}, "get,post:CreateTopic")
	beego.Router("/user/:user", &ctl.UserController{})
	beego.Router("/search", &ctl.SearchController{})
	beego.Router("/media/image/:size:string/:id:int", &ctl.MediaController{}, "get,post:Image")
	beego.AutoRouter(&ctl.ApiController{})

	topic := beego.NewNamespace("/t",
		beego.NSNamespace("/:topic",
			beego.NSRouter("/p/:post", &ctl.PostController{}),
			beego.NSRouter("/?:choice", &ctl.TopicController{}),
		),
	)
	beego.AddNamespace(topic)
}
