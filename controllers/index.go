package controllers

import (
	"log"
	"net/http"
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
	posts, err := getPostsForTopic(this.User, nil, choice)
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
