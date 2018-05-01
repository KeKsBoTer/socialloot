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
		// invalid image id
		c.Abort("404")
		return
	}
	size := c.Ctx.Input.Param(":size")
	if size != "small" && size != "original" {
		// invalid image size
		c.Abort("404")
		return
	}
	media := models.Media{
		Id: id,
	}
	if err := media.Read("Id"); err != nil {
		// image not found
		c.Abort("404")
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
