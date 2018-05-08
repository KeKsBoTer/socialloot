package controllers

import (
	"net/http"

	"github.com/KeKsBoTer/socialloot/models"
)

// SubmitController handles new posts and topics
type SubmitController struct {
	NeedsAuthController
}

// Submit serves submit post page
func (c *SubmitController) Submit() {
	c.TplName = "pages/submit/submit.tpl"
	c.Data["Title"] = "Submit to Socialloot"

	submitType := c.GetString("type")
	if submitType != "text" && submitType != "link" && submitType != "image" {
		submitType = "text"
	}
	c.Data["Type"] = submitType

	topicName := c.GetString("topic")
	if len(topicName) < 1 {
		return
	}
	topic := models.Topic{
		Name: topicName,
	}
	if err := topic.Read("name"); err == nil {
		c.Data["Topic"] = topic
	} else {
		c.Redirect(c.URLFor("SubmitController.Submit"), http.StatusSeeOther)
	}
}

// CreateTopic serves topic creation page
func (c *SubmitController) CreateTopic() {
	c.TplName = "pages/submit/createTopic.tpl"
	c.Data["Title"] = "Create a topic"
}
