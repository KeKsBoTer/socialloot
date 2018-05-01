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

func (c *UserController) Get() {
	choice := UserChoice(c.Ctx.Input.Param(":choice"))
	if len(choice) < 1 {
		choice = Posts
	}
	if !choice.IsValid() {
		c.CustomAbort(http.StatusBadRequest, "invalid choice")
		return
	}

	userName := c.Ctx.Input.Param(":user")
	user := &models.User{
		Name: userName,
	}
	if err := user.Read("Name"); err != nil {
		c.Abort("404")
		return
	}

	switch choice {
	case Posts:
		// get all the posts the user published
		if posts, err := getPostsForUser(user, c.User); err == nil {
			c.Data["Posts"] = posts
		} else {
			c.Abort("500")
			return
		}
	case Comments:
		// get all the posts the user published
		if comments, err := getCommentsForUser(user); err == nil {
			c.Data["Comments"] = comments
		} else {
			c.Abort("500")
			return
		}
	}

	c.Data["Choice"] = choice
	c.Data["User"] = user
	c.TplName = "pages/users/page.tpl"
}

// IsValid checks if choice is hot or new
func (c UserChoice) IsValid() bool {
	return c == Comments || c == Posts
}

func getPostsForUser(user, viewer *models.User) (*models.PostMetaDataList, error) {
	var posts models.PostList
	if _, err := models.Posts().Filter("user", user).OrderBy("-Date").RelatedSel().All(&posts); err != nil {
		return nil, err
	}
	metas := posts.ToMetaData()
	for _, p := range *metas {
		p.ReadVoteData(viewer)
	}
	return metas, nil
}

func getCommentsForUser(user *models.User) (*models.CommentMetaDataList, error) {
	var comments models.CommentList
	if _, err := models.Comments().Filter("user", user).OrderBy("-Date").RelatedSel().All(&comments); err != nil {
		return nil, err
	}
	metas := comments.ToMetaData()
	for _, c := range *metas {
		if err := c.ReadVoteData(user); err != nil {
			beego.Error("Cannot read user vote data on post:", err)
		}
	}
	return metas, nil
}
