package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/KeKsBoTer/socialloot/controllers:SearchController"] = append(beego.GlobalControllerRouter["github.com/KeKsBoTer/socialloot/controllers:SearchController"],
		beego.ControllerComments{
			Method: "Search",
			Router: `/search`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
