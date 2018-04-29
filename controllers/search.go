package controllers

import (
	"net/http"

	"github.com/astaxie/beego"

	"github.com/KeKsBoTer/socialloot/models"
	"github.com/astaxie/beego/orm"
)

type SearchController struct {
	AuthController
}

// SearchChoice is the type of content to search for
type SearchChoice string

const (
	SearchPosts  SearchChoice = "posts"
	SearchTopics SearchChoice = "topics"
	SearchUsers  SearchChoice = "users"
)

// IsValid checks if choice is hot or new
func (c SearchChoice) IsValid() bool {
	return c == SearchPosts || c == SearchTopics || c == SearchUsers
}

func (this *SearchController) Get() {
	query := this.GetString("query")
	choice := SearchChoice(this.GetString("choice"))

	if len(choice) < 1 {
		choice = SearchPosts
	}
	if !choice.IsValid() {
		this.CustomAbort(http.StatusBadRequest, "invalid choice")
		return
	}

	pageTitle := "Search"
	if len(query) > 0 {
		pageTitle += ": " + query
		var err error
		var result interface{}
		switch choice {
		case SearchPosts:
			if posts, err := searchPosts(query); err == nil {
				result = posts
			}
		case SearchTopics:
			if topics, err := searchTopics(query); err == nil {
				result = topics
			}
		case SearchUsers:
			if users, err := searchUsers(query); err == nil {
				result = users
			}
		}
		if err != nil {
			beego.Error("Error searching for ", choice, ". Query:", query)
			beego.Error(err)
		} else {
			this.Data["SearchResult"] = result
		}
	}
	this.Data["SearchQuery"] = query
	this.Data["Title"] = pageTitle
	this.Data["Choice"] = choice
	this.TplName = "pages/search.tpl"
}

func searchPosts(key string) (*[]*models.Post, error) {
	isNotImage := orm.NewCondition().Or("type", models.PostTypeText).Or("type", models.PostTypeLink)
	inText := orm.NewCondition().AndCond(isNotImage).And("content__icontains", key)
	cond := orm.NewCondition().Or("title__icontains", key).OrCond(inText)
	var posts []*models.Post
	if _, err := models.Posts().SetCond(cond).Limit(20).RelatedSel("topic").RelatedSel("user").All(&posts); err != nil {
		return nil, err
	}
	return &posts, nil
}

func searchTopics(key string) (*[]*models.Topic, error) {
	cond := orm.NewCondition().Or("title__icontains", key).Or("name__icontains", key).Or("description__icontains", key)
	var topics []*models.Topic
	if _, err := models.Topics().SetCond(cond).Limit(20).All(&topics); err != nil {
		return nil, err
	}
	return &topics, nil
}

func searchUsers(key string) (*[]*models.User, error) {
	var users []*models.User
	if _, err := models.Users().Filter("name__icontains", key).Limit(20).All(&users); err != nil {
		return nil, err
	}
	return &users, nil
}
