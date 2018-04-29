package controllers

import (
	"net/http"

	"github.com/astaxie/beego"

	"github.com/KeKsBoTer/socialloot/models"
)

type UserController struct {
	AuthController
}

// UserChoice is the kind of information that is displayes about the user
type UserChoice string

const (
	// Comments will display all comments the user wrote
	Comments UserChoice = "comments"
	// Posts will display all posts the user published
	Posts UserChoice = "posts"
)

func (this *UserController) Get() {
	choice := UserChoice(this.Ctx.Input.Param(":choice"))
	if len(choice) < 1 {
		choice = Posts
	}
	if !choice.IsValid() {
		this.CustomAbort(http.StatusBadRequest, "invalid choice")
		return
	}

	userName := this.Ctx.Input.Param(":user")
	user := &models.User{
		Name: userName,
	}
	if err := user.Read("Name"); err != nil {
		this.Abort("404")
		return
	}

	switch choice {
	case Posts:
		// get all the posts the user published
		if posts, err := getPostsForUser(user, this.User); err == nil {
			this.Data["Posts"] = posts
		} else {
			this.Abort("500")
			return
		}
	case Comments:
		// get all the posts the user published
		if comments, err := getCommentsForUser(user); err == nil {
			this.Data["Comments"] = comments
		} else {
			this.Abort("500")
			return
		}
	}

	this.Data["Choice"] = choice
	this.Data["User"] = user
	this.TplName = "pages/users/page.tpl"
}

// IsValid checks if choice is hot or new
func (c UserChoice) IsValid() bool {
	return c == Comments || c == Posts
}

func getPostsForUser(user, viewer *models.User) (*[]*models.Post, error) {
	var posts []*models.Post
	if _, err := models.Posts().Filter("user", user).OrderBy("-Date").RelatedSel().All(&posts); err != nil {
		return nil, err
	}
	for _, p := range posts {
		p.ReadVoteData(viewer)
	}
	return &posts, nil
}

func getCommentsForUser(user *models.User) (*[]*models.Comment, error) {
	var comments []*models.Comment
	if _, err := models.Comments().Filter("user", user).OrderBy("-Date").RelatedSel().All(&comments); err != nil {
		return nil, err
	}
	for _, c := range comments {
		if err := c.ReadVoteData(user); err != nil {
			beego.Error("Cannot read user vote data on post:", err)
		}
	}
	return &comments, nil
}
