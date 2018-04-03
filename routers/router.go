package routers

import (
	ctl "github.com/KeKsBoTer/socialloot/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &ctl.IndexController{})
	beego.Router("/login", &ctl.LoginController{}, "get:LoginPage;post:Login")
	beego.Router("/logout", &ctl.LoginController{}, "get:Logout")
	beego.Router("/signup", &ctl.LoginController{}, "get:SignupPage;post:Signup")
	beego.Router("/submit", &ctl.SubmitController{}, "get,post:Submit")
	beego.Router("/createtopic", &ctl.SubmitController{}, "get,post:CreateTopic")
	beego.Router("/user/:user", &ctl.UserController{})
	beego.Router("/media/image/:size:string/:id:int", &ctl.MediaController{}, "get,post:Image")
	beego.AutoRouter(&ctl.ApiController{})

	topic := beego.NewNamespace("/topic",
		beego.NSNamespace(":topic",
			beego.NSRouter("/", &ctl.TopicController{}),
			beego.NSRouter("/:post", &ctl.PostController{}),
		),
	)
	beego.AddNamespace(topic)
}
