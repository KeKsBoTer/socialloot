package controllers

import (
	"net/http"
	"strings"

	"github.com/astaxie/beego"

	"github.com/KeKsBoTer/socialloot/models"
	"github.com/astaxie/beego/orm"
)

// SearchController provides search page and outputs search results
type SearchController struct {
	AuthController
}

// Get serves page for search page
func (c *SearchController) Get() {
	query := strings.Trim(c.GetString("query"), " ")
	choice := SearchChoice(c.GetString("choice"))

	if len(choice) < 1 {
		choice = SearchPosts
	}
	if !choice.IsValid() {
		c.CustomAbort(http.StatusBadRequest, "invalid choice")
		return
	}

	pageTitle := "Search"
	if len(query) > 0 {
		pageTitle += ": " + query

		var err error
		var result interface{}
		switch choice {
		case SearchPosts:
			result, err = searchPosts(query, c.User)
		case SearchTopics:
			result, err = searchTopics(query)
		case SearchUsers:
			result, err = searchUsers(query)
		}
		if err == nil {
			c.Data["SearchResult"] = result
		} else {
			beego.Error("Error searching for ", choice, ". Query:", query, "\n", err)
		}
	}
	c.Data["SearchQuery"] = query
	c.Data["Title"] = pageTitle
	c.Data["Choice"] = choice
	c.TplName = "pages/search.tpl"
}

// find all posts with the given key in title, link or text
func searchPosts(key string, viewer *models.User) (*models.PostMetaDataList, error) {
	// check if post is of type text or link
	isNotImage := orm.NewCondition().Or("type", models.PostTypeText).Or("type", models.PostTypeLink)
	// check if search query is in text or link
	// no searching in images
	inText := orm.NewCondition().AndCond(isNotImage).And("content__icontains", key)
	// check if query is in title or text
	cond := orm.NewCondition().Or("title__icontains", key).OrCond(inText)

	var posts models.PostList
	if _, err := models.Posts().SetCond(cond).Limit(20).RelatedSel("topic").RelatedSel("user").All(&posts); err != nil {
		return nil, err
	}
	meta := posts.ToMetaData()
	meta.ReadVoteData(viewer)
	return meta, nil
}

// find all topics with the given key in title, name or description
func searchTopics(key string) (*[]*models.Topic, error) {
	cond := orm.NewCondition().Or("title__icontains", key).Or("name__icontains", key).Or("description__icontains", key)
	var topics []*models.Topic
	if _, err := models.Topics().SetCond(cond).Limit(20).All(&topics); err != nil {
		return nil, err
	}
	return &topics, nil
}

// find all users with the given key in the name
func searchUsers(key string) (*[]*models.User, error) {
	var users []*models.User
	if _, err := models.Users().Filter("name__icontains", key).Limit(20).All(&users); err != nil {
		return nil, err
	}
	return &users, nil
}

// SearchChoice is the type of content to search for
type SearchChoice string

const (
	// SearchPosts search for posts
	SearchPosts SearchChoice = "posts"
	//SearchTopics search for topics
	SearchTopics SearchChoice = "topics"
	// SearchUsers search for users
	SearchUsers SearchChoice = "users"
)

// IsValid checks if choice is hot or new
func (c SearchChoice) IsValid() bool {
	return c == SearchPosts || c == SearchTopics || c == SearchUsers
}
