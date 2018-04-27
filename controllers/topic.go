package controllers

import (
	"log"
	"net/http"

	"github.com/KeKsBoTer/socialloot/lib"
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
	posts, err := getPostsForTopic(this.User, topic, choice)
	if err != nil {
		log.Println(err)
		this.Abort("505")
		return
	}
	this.Data["Posts"] = posts
	this.Layout = "base.tpl"
	this.TplName = "pages/posts/topic.tpl"
}

func getPostsForTopic(user *models.User, topic *models.Topic, choice Choice) (*[]*models.Post, error) {
	var posts []*models.Post
	all := models.Posts()
	if topic != nil {
		all = all.Filter("topic", topic.Id)
	}
	// hot ranking algorithm is heavily based on date, so ordering by date on select will speed up sorting
	if _, err := all.OrderBy("-Date").RelatedSel().All(&posts); err != nil {
		return nil, err
	}
	for _, p := range posts {
		p.ReadVoteData(user)
	}
	switch choice {
	case New:
		// allready sorted
	case Hot:
		lib.SortByRank(posts)
	}
	return &posts, nil
}

// IsValid checks if choice is hot or new
func (c Choice) IsValid() bool {
	return c == Hot || c == New
}
