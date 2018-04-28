package controllers

import (
	"github.com/KeKsBoTer/socialloot/models"
	"github.com/astaxie/beego/orm"
)

type SearchController struct {
	AuthController
}

func (this *SearchController) Get() {
	query := this.GetString("query")

	pageTitle := "Search"
	if len(query) > 0 {
		pageTitle += ": " + query
		if posts, err := search(query); err == nil && len(*posts) > 0 {
			this.Data["Posts"] = posts
		}
	}
	this.Data["SearchQuery"] = query
	this.Data["Title"] = pageTitle
	this.TplName = "pages/search.tpl"
}

func search(key string) (*[]*models.Post, error) {
	inText := orm.NewCondition().And("type", models.PostTypeText).And("content__icontains", key)
	cond := orm.NewCondition().Or("title__icontains", key).OrCond(inText)
	var posts []*models.Post
	if _, err := models.Posts().SetCond(cond).Limit(20).RelatedSel("topic").All(&posts); err != nil {
		return nil, err
	}
	return &posts, nil
}
