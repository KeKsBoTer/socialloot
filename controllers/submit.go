package controllers

import (
	"html/template"
	"net/http"

	"github.com/KeKsBoTer/socialloot/lib"
	"github.com/KeKsBoTer/socialloot/models"
	"github.com/astaxie/beego"
)

type SubmitController struct {
	BaseController
}

func (c *SubmitController) Submit() {
	c.TplName = "submit/submit.tpl"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())

	if !c.Ctx.Input.IsPost() {
		return
	}

	var err error
	flash := beego.NewFlash()

	p := &models.Post{}
	if err = c.ParseForm(p); err != nil {
		flash.Error("Submit invalid!")
		flash.Store(&c.Controller)
		return
	}
	p.User = c.GetLogin()
	if err = models.IsValid(p); err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		return
	}

	err = lib.Submit(p)
	if err != nil {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		return
	}

	flash.Success("Post submited")
	flash.Store(&c.Controller)

	c.Redirect(c.URLFor("UsersController.Index"), http.StatusSeeOther)
}

func (c *SubmitController) CreateTopic() {
	c.TplName = "submit/createTopic.tpl"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())

	if !c.Ctx.Input.IsPost() {
		return
	}

	var err error
	flash := beego.NewFlash()

	t := &models.Topic{}
	if err = c.ParseForm(t); err != nil {
		flash.Error("Submit invalid!")
		flash.Store(&c.Controller)
		return
	}
	if err = models.IsValid(t); err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		return
	}

	err = lib.CreateTopic(t)
	if err != nil {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		return
	}

	flash.Success("Topic created")
	flash.Store(&c.Controller)

	c.Redirect(c.URLFor("UsersController.Index"), http.StatusSeeOther)
}
