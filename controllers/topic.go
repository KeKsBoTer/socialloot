package controllers

import (
	"log"

	"github.com/KeKsBoTer/socialloot/lib"
	"github.com/KeKsBoTer/socialloot/models"
)

// TopicController serves topic page
type TopicController struct {
	AuthController
}

// Get serves topic pages
func (c *TopicController) Get() {
	topicName := c.Ctx.Input.Param(":topic")
	choice := Choice(c.Ctx.Input.Param(":choice"))
	if len(choice) < 1 {
		choice = Hot
	}
	if !choice.IsValid() {
		c.Abort("404")
		return
	}
	c.Data["Choice"] = choice
	topic := &models.Topic{
		Name: topicName,
	}
	if err := topic.Read("name"); err != nil {
		c.Abort("404")
		return
	}
	c.Data["Topic"] = topic
	posts, err := getPostsForTopic(c.User, topic, choice)
	if err != nil {
		log.Println(err)
		c.Abort("500")
		return
	}
	c.Data["Posts"] = posts
	c.Layout = "base.tpl"
	c.TplName = "pages/posts/topic.tpl"
}

// loads all posts for a topic and orders them by the given choice:
//  - new: newest first
//  - hot: posts with hightest rank first (see lib.SortByRank)
// also the vote data for the viewer is loaded
func getPostsForTopic(viewer *models.User, topic *models.Topic, choice Choice) (*models.PostMetaDataList, error) {
	var posts models.PostList
	all := models.Posts()
	if topic != nil {
		all = all.Filter("topic", topic.Id)
	}
	// hot ranking algorithm is heavily based on date, so ordering by date on select will speed up sorting
	if _, err := all.OrderBy("-Date").RelatedSel().All(&posts); err != nil {
		return nil, err
	}
	metas := posts.ToMetaData()
	for _, m := range *metas {
		m.ReadVoteData(viewer)
	}
	switch choice {
	case New:
		// allready sorted
	case Hot:
		lib.SortByRank(*metas)
	}
	return metas, nil
}

// Choice defines the way the listed posts are sorted
type Choice string

const (
	// Hot means posts with most votes first
	Hot Choice = "hot"
	// New means newest posts first
	New Choice = "new"
)

// IsValid checks if choice is hot or new
func (c Choice) IsValid() bool {
	return c == Hot || c == New
}
