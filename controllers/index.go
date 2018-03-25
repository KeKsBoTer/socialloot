package controllers

import (
	"github.com/KeKsBoTer/socialloot/models"
)

type IndexController struct {
	AuthController
}

func (this *IndexController) Get() {
	this.Data["Title"] = "Socailloot: like reddit but different"
	this.TplName = "index.tpl"
	var topics []*models.Topic
	if _, err := models.Topics().All(&topics); err != nil {
		this.Abort("505")
		return
	}
	this.Data["Topics"] = topics
}
