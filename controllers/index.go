package controllers

type IndexController struct {
	AuthController
}

func (this *IndexController) Get() {
	this.Data["Title"] = "Socailloot: like reddit but different"
	this.TplName = "pages/index.tpl"
}
