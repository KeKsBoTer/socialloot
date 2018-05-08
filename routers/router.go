// @APIVersion 1.0.0
// @Title socialloot webpage
// @Description socialloot webpage

package routers

import (
	ctl "github.com/KeKsBoTer/socialloot/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// custom error page renderer
	beego.ErrorController(&ctl.ErrorController{})

	// auth pages
	beego.Router("/login", &ctl.LoginController{}, "get:LoginPage;post:Login")
	beego.Router("/logout", &ctl.LoginController{}, "get:Logout")
	beego.Router("/signup", &ctl.LoginController{}, "get:SignupPage;post:Signup")

	beego.Router("/submit", &ctl.SubmitController{}, "get,post:Submit")
	beego.Router("/createtopic", &ctl.SubmitController{}, "get,post:CreateTopic")

	// post and topic pages
	beego.AddNamespace(
		beego.NewNamespace("/t",
			beego.NSNamespace("/:topic",
				beego.NSRouter("/p/:post", &ctl.PostController{}),
				beego.NSRouter("/?:choice", &ctl.TopicController{}),
			),
		),
	)

	// user profile apge
	beego.Router("/user/:user/?:choice", &ctl.UserController{})
	beego.Router("/search", &ctl.SearchController{})
	beego.Router("/media/image/:size:string/:id:int", &ctl.ImageController{}, "get,post:Image")
	api := &ctl.APIController{}
	beego.AddNamespace(
		beego.NewNamespace("/api",
			beego.NSRouter("/submit", api, "post:Submit"),
			beego.NSRouter("/createtopic", api, "post:CreateTopic"),
			beego.NSRouter("/comment", api, "post:Comment"),
			beego.NSRouter("/vote", api, "post:Vote"),
			beego.NSRouter("/delete", api, "post:Delete"),
		),
	)

	beego.Router("/?:choice(hot|new)", &ctl.IndexController{})
	// since a post id is unique we can redirect /post-id to the post
	beego.Router("/?:post", &ctl.PostController{}, "get:Redirect")
}
