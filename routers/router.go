package routers

import (
	ctl "github.com/KeKsBoTer/socialloot/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/?:choice(hot|new)", &ctl.IndexController{})
	beego.Router("/?:post", &ctl.PostController{}, "get:Redirect")
	beego.Router("/login", &ctl.LoginController{}, "get:LoginPage;post:Login")
	beego.Router("/logout", &ctl.LoginController{}, "get:Logout")
	beego.Router("/signup", &ctl.LoginController{}, "get:SignupPage;post:Signup")
	beego.Router("/submit", &ctl.SubmitController{}, "get,post:Submit")
	beego.Router("/createtopic", &ctl.SubmitController{}, "get,post:CreateTopic")
	beego.Router("/user/:user/?:choice", &ctl.UserController{})
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
