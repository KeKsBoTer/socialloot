package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/KeKsBoTer/socialloot/models"
	"github.com/astaxie/beego/orm"
)

type IndexController struct {
	AuthController
}

func (this *IndexController) Get() {

	choice := Choice(this.Ctx.Input.Param(":choice"))
	if len(choice) < 1 {
		choice = Hot
	}
	if !choice.IsValid() {
		this.CustomAbort(http.StatusBadRequest, "invalid choice")
		return
	}
	posts, err := getPosts(this.User, choice)
	if err != nil {
		log.Println(err)
		this.Abort("505")
		return
	}
	this.Data["Posts"] = posts
	this.Data["Choice"] = choice

	this.Data["Title"] = "Socailloot: like reddit but different"
	this.TplName = "pages/index.tpl"
}

func getPosts(user *models.User, choice Choice) (*[]*models.Post, error) {
	var posts []*models.Post
	switch choice {
	case New:
		if _, err := models.Posts().OrderBy("-Date").RelatedSel().All(&posts); err != nil {
			return nil, err
		}
	case Hot:
		_, err := orm.NewOrm().Raw(`
			WITH votes AS(
				SELECT item,sum(action) as votes
				FROM vote
				GROUP BY item
			)
			SELECT p.*, ifnull(v.votes,0) as votes
			FROM post p
			LEFT OUTER JOIN votes v on (p.id = v.item)
			ORDER BY ifnull(v.votes,0) desc`).QueryRows(&posts)
		for _, p := range posts {
			o := orm.NewOrm()
			o.LoadRelated(p, "topic")
			o.LoadRelated(p, "user")
		}
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("invalid choice")
	}
	for _, p := range posts {
		p.ReadVoteData(user)
	}
	return &posts, nil
}
