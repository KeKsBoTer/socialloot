package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/astaxie/beego/orm"

	"github.com/KeKsBoTer/socialloot/models"
)

// Choice defines the way the listed posts are sorted
type Choice string

const (
	// Hot means posts with most votes first
	Hot Choice = "hot"
	// New means newest posts first
	New Choice = "new"
)

type TopicController struct {
	AuthController
}

func (this *TopicController) Get() {
	topicName := this.Ctx.Input.Param(":topic")
	choice := Choice(this.Ctx.Input.Param(":choice"))
	if len(choice) < 1 {
		choice = Hot
	}
	if !choice.IsValid() {
		this.CustomAbort(http.StatusBadRequest, "invalid choice")
		return
	}
	this.Data["Choice"] = choice
	topic := &models.Topic{
		Name: topicName,
	}
	if err := topic.Read("name"); err != nil {
		this.Abort("404")
		return
	}
	this.Data["Topic"] = topic
	posts, err := getPostsForTopic(topic, this.User, choice)
	if err != nil {
		log.Println(err)
		this.Abort("505")
		return
	}
	this.Data["Posts"] = posts
	this.Layout = "base.tpl"
	this.TplName = "pages/posts/topic.tpl"
}

func getPostsForTopic(topic *models.Topic, user *models.User, choice Choice) (*[]*models.Post, error) {
	var posts []*models.Post
	switch choice {
	case New:
		if _, err := models.Posts().Filter("Topic", topic.Id).OrderBy("-Date").RelatedSel().All(&posts); err != nil {
			return nil, err
		}
	case Hot:
		_, err := orm.NewOrm().Raw(`
			WITH votes AS(
				SELECT item,sum(action) as votes
				FROM vote
				GROUP BY item
			)
			SELECT p.*, ifnull(v.votes,0) as votes
			FROM post p
			LEFT OUTER JOIN votes v on (p.id = v.item)
			WHERE p.topic_id = ?
			ORDER BY ifnull(v.votes,0) desc`, topic.Id).QueryRows(&posts)
		for _, p := range posts {
			o := orm.NewOrm()
			o.LoadRelated(p, "topic")
			o.LoadRelated(p, "user")
		}
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("invalid choice")
	}
	for _, p := range posts {
		p.ReadVoteData(user)
	}
	return &posts, nil
}

// IsValid checks if choice is hot or new
func (c Choice) IsValid() bool {
	return c == Hot || c == New
}
