package controllers

import (
	"html/template"
)

type SubmitController struct {
	NeedsAuthController
}

func (c *SubmitController) Submit() {
	c.TplName = "submit/submit.tpl"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["Title"] = "Submit to Socialloot"
}

func (c *SubmitController) CreateTopic() {
	c.TplName = "submit/createTopic.tpl"
	c.Data["Title"] = "Create a topic"
}
