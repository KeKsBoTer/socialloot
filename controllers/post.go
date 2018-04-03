package controllers

import (
	"github.com/KeKsBoTer/socialloot/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
	if _, err := orm.NewOrm().LoadRelated(&post, "user"); err != nil {
		this.Abort("505")
		return
	}

	// loading comments
	if err := post.ReadComments(this.User); err != nil {
		beego.Error(err)
	}

	// loading total votes count and user vote
	if err := post.ReadVoteData(this.User); err != nil {
		beego.Error(err)
	}

	this.Data["Post"] = post
	this.Layout = "base.tpl"
	this.TplName = "pages/posts/post.tpl"
}
