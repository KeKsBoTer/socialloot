package controllers

import (
	"net/http"

	"github.com/KeKsBoTer/socialloot/models"
)

type SubmitController struct {
	NeedsAuthController
}

func (c *SubmitController) Submit() {
	c.TplName = "pages/submit/submit.tpl"
	c.Data["Title"] = "Submit to Socialloot"
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

func (c *SubmitController) CreateTopic() {
	c.TplName = "pages/submit/createTopic.tpl"
	c.Data["Title"] = "Create a topic"
}
