package controllers

import (
	"net/http"

	"github.com/KeKsBoTer/socialloot/lib"
	"github.com/KeKsBoTer/socialloot/models"
	"github.com/astaxie/beego"
)

type PostController struct {
	AuthController
}

func (c *PostController) Get() {
	topicName := c.Ctx.Input.Param(":topic")
	topic, err := models.ReadTopic(topicName)
	if err != nil {
		// topic not found
		c.Abort("404")
		return
	}

	postID := c.Ctx.Input.Param(":post")
	post, err := models.ReadPost(postID, true)
	if err != nil {
		// post not found
		c.Abort("404")
		return
	}
	meta := post.NewMetaData()
	// loading comments
	if err := meta.ReadComments(c.User); err != nil {
		beego.Error(err)
	}

	// loading total votes count and user vote
	if err := meta.ReadVoteData(c.User); err != nil {
		beego.Error(err)
	}

	// check is user created the post, so he can delete it
	c.Data["CanDelete"] = c.User != nil && post.User.Id == c.User.Id
	c.Data["Topic"] = topic
	c.Data["Post"] = meta

	c.Layout = "base.tpl"
	c.TplName = "pages/posts/post.tpl"
}

func (c *PostController) Redirect() {
	id := c.Ctx.Input.Param(":post")
	post := models.Post{
		Id: id,
	}
	if err := post.Read(true); err == nil {
		url := lib.URLForItem(post)
		if len(url) > 0 {
			c.Ctx.Redirect(http.StatusTemporaryRedirect, url)
		} else {
			// invalid post item
			c.Abort("500")
		}
		return
	}
	c.Abort("404")
}
