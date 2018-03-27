package controllers

import (
	"log"

	"github.com/astaxie/beego"

	"github.com/KeKsBoTer/socialloot/lib"
	"github.com/KeKsBoTer/socialloot/models"
)

type ApiController struct {
	NeedsAuthController
}

func (c *ApiController) Vote() {
	dir, err := c.GetInt("dir")
	if err != nil {
		c.Abort("400")
		return
	}
	id := c.GetString("id")
	if len(id) < 1 {
		c.Abort("400")
		return
	}
	if err := lib.VoteOnPost(models.UserVote(dir), id, c.GetLogin()); err != nil {
		log.Println(err)
		c.Abort("500")
		return
	}
	c.Data["json"] = "success"
	c.ServeJSON(true)
}

type ApiResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Dest    string `json:"dest"`
}

func (c *ApiController) Submit() {
	// server answer as json
	r := ApiResponse{}
	c.Data["json"] = &r
	defer c.ServeJSON(true)

	p := &models.Post{}
	if err := c.ParseForm(p); err != nil {
		r.Success = false
		r.Message = err.Error()
		return
	}
	p.User = c.User
	if err := models.IsValid(p); err != nil {
		r.Success = false
		r.Message = err.Error()
		return
	}
	err := lib.Submit(p)
	if err != nil {
		r.Success = false
		r.Message = err.Error()
		return
	}
	r.Success = true
	r.Dest = lib.URLForItem(*p)
}

func (c *ApiController) CreateTopic() {
	// server answer as json
	r := ApiResponse{}
	c.Data["json"] = &r
	defer c.ServeJSON(true)

	t := &models.Topic{}
	if err := c.ParseForm(t); err != nil {
		r.Success = false
		r.Message = err.Error()
		return
	}
	if err := models.IsValid(t); err != nil {
		r.Success = false
		r.Message = err.Error()
		return
	}
	err := lib.CreateTopic(t)
	if err != nil {
		r.Success = false
		r.Message = err.Error()
		return
	}
	r.Success = true
	r.Dest = lib.URLForItem(*t)
}

func apiResponse(c *beego.Controller) *ApiResponse {
	r := ApiResponse{}
	if dst := c.GetString("dest"); len(dst) > 0 {
		r.Dest = dst
	}
	c.Data["json"] = &r
	return &r
}
