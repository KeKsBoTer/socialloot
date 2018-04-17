package controllers

import (
	"log"

	"github.com/KeKsBoTer/socialloot/models"
)

type TopicController struct {
	AuthController
}

func (this *TopicController) Get() {
	topicName := this.Ctx.Input.Param(":topic")
	topic := &models.Topic{
		Name: topicName,
	}
	if err := topic.Read("name"); err != nil {
		this.Abort("404")
		return
	}
	this.Data["Topic"] = topic
	posts, err := getPostsForTopic(topic, this)
	if err != nil {
		log.Println(err)
		this.Abort("505")
		return
	}
	this.Data["Posts"] = posts
	this.Layout = "base.tpl"
	this.TplName = "pages/posts/topic.tpl"
}

func getPostsForTopic(topic *models.Topic, c *TopicController) (*[]*models.Post, error) {
	var posts []*models.Post
	if _, err := models.Posts().Filter("Topic", topic.Id).OrderBy("-Date").RelatedSel().All(&posts); err != nil {
		return nil, err
	}
	for i := range posts {
		posts[i].ReadVoteData(c.User)
	}
	return &posts, nil
}
