package controllers

import (
	"errors"

	"github.com/astaxie/beego"

	"github.com/KeKsBoTer/socialloot/lib"
	"github.com/KeKsBoTer/socialloot/models"
)

type ApiController struct {
	NeedsAuthController
}

func (c *ApiController) Vote() {
	form := &models.VoteForm{}
	handleForm(form, &c.AuthController, func(r *ApiResponse) {
		err := lib.VoteOnItem(form.Direction, form.Item, c.User)
		if err != nil {
			r.Fail(err)
			return
		}
		r.Success = true
	})
}

// Size constants
const (
	MB = 1 << 20
)

func (c *ApiController) Submit() {
	form := &models.SubmitForm{}
	handleForm(form, &c.AuthController, func(r *ApiResponse) {
		topic := models.Topic{
			Name: form.TopicName,
		}
		if err := topic.Read("name"); err != nil {
			r.Fail(errors.New("Topic not found"))
			return
		}
		if form.Type == models.PostTypeImage {
			file, header, err := c.GetFile("content")
			if err != nil {
				r.Fail(err)
				return
			}
			if header.Size > 20*MB {
				r.Fail(errors.New("Maximum file size is 20MB"))
				return
			}
			img := make([]byte, header.Size)
			if _, err := file.Read(img); err != nil {
				r.Fail(err)
				return
			}
			form.Content = string(img)
		}
		post, err := lib.Submit(form.Title, form.Content, form.Type, &topic, c.User)
		if err != nil {
			r.Fail(err)
			return
		}
		r.Success = true
		r.Dest = lib.URLForItem(post)
	})
}

func (c *ApiController) CreateTopic() {
	form := &models.CreateTopicForm{}
	handleForm(form, &c.AuthController, func(r *ApiResponse) {
		topic, err := lib.CreateTopic(form.Name, form.Title, form.Description)
		if err != nil {
			r.Fail(err)
			return
		}
		r.Success = true
		r.Dest = lib.URLForItem(topic)
	})
}

func (c *ApiController) Comment() {
	form := &models.CommentForm{}
	handleForm(form, &c.AuthController, func(r *ApiResponse) {
		if err := lib.CommentOnPost(form.Comment, form.Item, c.User); err != nil {
			r.Fail(err)
			return
		}
		r.Success = true
	})
}

func (c *ApiController) Delete() {
	form := &models.DeleteForm{}
	handleForm(form, &c.AuthController, func(r *ApiResponse) {
		if err := lib.DeletePost(form.Item, c.User); err != nil {
			r.Fail(err)
			return
		}
		r.Success = true
		r.Dest = "/"
	})
}

type ApiResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Dest    string `json:"dest"`
}

func NewApiResponse(c *beego.Controller) *ApiResponse {
	r := ApiResponse{}
	if dst := c.GetString("dest"); len(dst) > 0 {
		r.Dest = dst
	}
	c.Data["json"] = &r
	return &r
}

func (a *ApiResponse) Fail(err error) {
	a.Success = false
	a.Message = err.Error()
}

type FormHandlerFunc func(r *ApiResponse)

func handleForm(form interface{}, c *AuthController, f FormHandlerFunc) {
	r := NewApiResponse(&c.Controller)
	defer c.ServeJSON(true)
	if err := c.ParseForm(form); err != nil {
		r.Fail(err)
		return
	}
	if err := models.IsValid(form); err != nil {
		r.Fail(err)
		return
	}
	f(r)
}
