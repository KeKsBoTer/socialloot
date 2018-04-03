package controllers

import (
	"strconv"

	"github.com/KeKsBoTer/socialloot/models"
)

type MediaController struct {
	AuthController
}

func (c *MediaController) Image() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.CustomAbort(400, "id musst be number")
		return
	}
	size := c.Ctx.Input.Param(":size")
	if size != "small" && size != "original" {
		c.CustomAbort(400, "invalid image size (small or original)")
		return
	}
	media := models.Media{
		Id: id,
	}
	if err := media.Read("Id"); err != nil {
		c.CustomAbort(400, "image does not exist")
		return
	}
	switch size {
	case "small":
		c.Ctx.Output.Body([]byte(media.Thumbnail))
	case "original":
		c.Ctx.Output.Body([]byte(media.File))
	}
	c.Ctx.Output.ContentType("image/*")
}
