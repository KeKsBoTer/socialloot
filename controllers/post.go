package controllers

import (
	"github.com/KeKsBoTer/socialloot/models"
)

type PostController struct {
	AuthController
}

func (this *PostController) Get() {
	topicName := this.Ctx.Input.Param(":topic")
	topic := models.Topic{
		Name: topicName,
	}
	if err := topic.Read("Name"); err != nil {
		this.Abort("404")
		return
	}
	this.Data["Topic"] = topic

	postID := this.Ctx.Input.Param(":post")
	post := models.Post{
		Id: postID,
	}
	if err := post.Read(); err != nil {
		this.Abort("404")
		return
	}
	post.ReadVoteData(this.User)
	this.Data["Post"] = post
	this.Layout = "base.tpl"
	this.TplName = "pages/posts/post.tpl"
}
