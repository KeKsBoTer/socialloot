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
	topic := &models.Topic{}
	if err := models.Topics().Filter("name", topicName).One(topic); err != nil {
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
	this.TplName = "posts/topic.tpl"
}

func getPostsForTopic(topic *models.Topic, c *TopicController) (*[]*models.Post, error) {
	var posts []*models.Post
	if _, err := models.Posts().Filter("Topic", topic.Id).RelatedSel().All(&posts); err != nil {
		return nil, err
	}
	for i, p := range posts {
		if c.User != nil {
			posts[i].VoteDir = models.GetUserVoteOnPost(c.User, p)
		}
		votes, err := models.GetVotesForPost(p)
		if err != nil {
			log.Println(err)
		} else {
			posts[i].Votes = votes
		}
	}
	return &posts, nil
}
