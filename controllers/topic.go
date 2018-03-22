package controllers

import (

	"github.com/KeKsBoTer/socialloot/models"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) PostsList() {
	topicName := this.Ctx.Input.Param(":topic")
	topic := &models.Topic{}
	if err := models.Topics().Filter("name", topicName).One(topic); err != nil {
		this.Abort("404")
		return
	}
	this.Data["Topic"] = topicName
	this.Layout = "base.tpl"
	this.TplName = "posts/topic.tpl"
}
