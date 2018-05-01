package controllers

import (
	"errors"
	"reflect"

	"github.com/astaxie/beego/validation"

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
			r.Fail("", err)
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
		topic, err := models.ReadTopic(form.TopicName)
		if err != nil {
			r.Fail("topic", errors.New("Topic not found"))
			return
		}
		if form.Type == models.PostTypeImage {
			file, header, err := c.GetFile("content")
			if err != nil {
				r.Fail("content", err)
				return
			}
			if header.Size > 20*MB {
				r.Fail("content", errors.New("Maximum file size is 20MB"))
				return
			}
			img := make([]byte, header.Size)
			if _, err := file.Read(img); err != nil {
				r.Fail("content", errors.New("cannot read image"))
				return
			}
			form.Content = string(img)
		}
		post, err := lib.Submit(form.Title, form.Content, form.Type, topic, c.User)
		if err != nil {
			r.Fail("", err)
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
			r.Fail("", err)
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
			r.Fail("", err)
			return
		}
		r.Success = true
	})
}

func (c *ApiController) Delete() {
	form := &models.DeleteForm{}
	handleForm(form, &c.AuthController, func(r *ApiResponse) {
		if err := lib.DeletePost(form.Item, c.User); err != nil {
			r.Fail("", err)
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
	Field   string `json:"field"`
}

func NewApiResponse(c *beego.Controller) *ApiResponse {
	r := ApiResponse{}
	if dst := c.GetString("dest"); len(dst) > 0 {
		r.Dest = dst
	}
	c.Data["json"] = &r
	return &r
}

func (a *ApiResponse) Fail(field string, err error) {
	a.Success = false
	a.Field = field
	a.Message = err.Error()
}

type FormHandlerFunc func(r *ApiResponse)

func handleForm(form interface{}, c *AuthController, f FormHandlerFunc) {
	r := NewApiResponse(&c.Controller)
	defer c.ServeJSON(true)
	if err := c.ParseForm(form); err != nil {
		r.Fail("", err)
		return
	}
	valid := validation.Validation{}
	b, err := valid.Valid(form)
	if err != nil {
		r.Fail("", errors.New("unexpected error"))
		return
	}
	if !b {
		// output first error
		first := valid.Errors[0]
		elm := reflect.TypeOf(form)
		// get struct not pointer
		if elm.Kind() == reflect.Ptr {
			elm = elm.Elem()
		}
		var formField string
		if field, found := elm.FieldByName(first.Field); found {
			// get field name from tag
			formField = field.Tag.Get("form")
		} else {
			// use struct field name
			formField = first.Field
		}
		r.Fail(formField, first)
		return
	}
	// handle form
	f(r)
}
