package controllers

import (
	"errors"
	"reflect"

	"github.com/astaxie/beego/validation"

	"github.com/astaxie/beego"

	"github.com/KeKsBoTer/socialloot/lib"
	"github.com/KeKsBoTer/socialloot/models"
)

// This is the socialloot api.
// It can only be accessed via HTTP POST by authenticated users.
// Most of the functions just validate the request and use functions from the lib package
// to handle the request.
//
// Functions:
// 	- vote on Post/Comment
// 	- submit or delete a post
//	- create a topic
// 	- publish a comment

// APIController is the controller for all api requests
// It needs authentication, otherwise an error will be returned
type APIController struct {
	NeedsAuthController
}

// Vote can be used to vote on a post
// See models.VoteForm for a closer description of the request format.
// This function uses lib.VoteOnItem(...)
func (c *APIController) Vote() {
	form := &models.VoteForm{}
	handleForm(form, &c.AuthController, func(r *APIResponse) {
		err := lib.VoteOnItem(form.Direction, form.Item, c.User)
		if err != nil {
			r.Fail("", err)
			return
		}
		r.Success = true
	})
}

// Size constants for 1 mega-byte
const (
	MB = 1 << 20
)

// Submit handles a submit submit request based on the content type.
// If the content type is image, it checks if the file is smaller than 20 MB.
// After successfully submitting the post, the user is redirected to the post.
// This function uses lib.Submit(..).
// For a detailes description of the request fields see models.SubmitForm.
func (c *APIController) Submit() {
	form := &models.SubmitForm{}
	handleForm(form, &c.AuthController, func(r *APIResponse) {
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
		// redirect to created post
		r.Dest = lib.URLForItem(post)
	})
}

// CreateTopic creates a new topic
// See models.CreateTopicForm
// This function uses lib.CreateTopic(...)
func (c *APIController) CreateTopic() {
	form := &models.CreateTopicForm{}
	handleForm(form, &c.AuthController, func(r *APIResponse) {
		topic, err := lib.CreateTopic(form.Name, form.Title, form.Description)
		if err != nil {
			r.Fail("", err)
			return
		}
		r.Success = true
		r.Dest = lib.URLForItem(topic)
	})
}

// Comment on a post with this function
// See models.CommentForm
// This function uses lib.CommentOnPost(...)
func (c *APIController) Comment() {
	form := &models.CommentForm{}
	handleForm(form, &c.AuthController, func(r *APIResponse) {
		if err := lib.CommentOnPost(form.Comment, form.Item, c.User); err != nil {
			r.Fail("", err)
			return
		}
		r.Success = true
	})
}

// Delete an existing post with this function
// See models.DeleteForm
// This function uses lib.DeletePost
func (c *APIController) Delete() {
	form := &models.DeleteForm{}
	handleForm(form, &c.AuthController, func(r *APIResponse) {
		if err := lib.DeletePost(form.Item, c.User); err != nil {
			r.Fail("", err)
			return
		}
		r.Success = true
		r.Dest = "/"
	})
}

// APIResponse is the model for all responses from the api
// The answer is converted in JSON format and sent to the client.
// example:
// {
//   "success": false,
//   "message": "unauthorized",
//   "dest": "/login",
//   "Field": "",
// }
type APIResponse struct {
	// Success tells the client with the request could be handled successfully
	// The JSON key is "success"
	Success bool `json:"success"`

	// Message describes the occurred error
	// If no error occurred the field is empty
	// The JSON key is "message"
	Message string `json:"message"`

	// Dest is the location the user is redirected to
	// This redirect is handled with JavaScript by the client
	// The JSON key is "dest"
	Dest string `json:"dest"`

	// Field is the key of the request parameter, which cased the error
	// This can be empty if no single field caused the error (e.g. unauthorized request)
	// The JSON key is "field"
	Field string `json:"field"`
}

// NewAPIResponse creates a new API response object from a controller
// If a destionation (a url) is provided in the HTTP GET parameters,
// the value is copied into the Dest field of the APIResponse.
// Also the response is applied to the controllers json field (c.Data["Json"])
// to tell beego, that the response is a JSON string.
func NewAPIResponse(c *beego.Controller) *APIResponse {
	r := APIResponse{}
	if dst := c.GetString("dest"); len(dst) > 0 {
		r.Dest = dst
	}
	// tell beego that it should write this object as json string to the response body
	c.Data["json"] = &r
	return &r
}

// Fail sets the responses success field to false
// and sets the message and field.
// The given error is converted to a string and used as message.
func (a *APIResponse) Fail(field string, err error) {
	a.Success = false
	a.Field = field
	a.Message = err.Error()
}

// FormHandlerFunc is a function that handles a parsed form
// The result of the request should be written to the api response parameter
// in this function.
type FormHandlerFunc func(r *APIResponse)

// This function parses the HTTP request and sets the corresponding fields in the given form
// The request is valideted using beego's validation (specificated in the model of the form as tag)
// If an error occurred during the validation, the error message and field is written to the response.
// At the end, if no error occurred, the given FormHandlerFunc is called.
func handleForm(form interface{}, c *AuthController, f FormHandlerFunc) {
	r := NewAPIResponse(&c.Controller)
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
