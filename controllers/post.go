package controllers

import (
	"net/http"

	"github.com/KeKsBoTer/socialloot/lib"
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
	this.Data["CanDelete"] = this.User != nil && post.User.Id == this.User.Id
	this.Layout = "base.tpl"
	this.TplName = "pages/posts/post.tpl"
}

func (this *PostController) Redirect() {
	id := this.Ctx.Input.Param(":post")
	if len(id) == models.ItemIDLength {
		post := models.Post{
			Id: id,
		}
		if err := post.Read(); err == nil {
			orm.NewOrm().LoadRelated(&post, "topic")
			this.Ctx.Redirect(http.StatusTemporaryRedirect, lib.URLForItem(post))
			return
		}
	}
	this.Abort("404")
}
