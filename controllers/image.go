package controllers

import (
	"strconv"

	"github.com/KeKsBoTer/socialloot/models"
)

// ImageController serves images via http
type ImageController struct {
	AuthController
}

// Image handles image request and writes the image data to the response body.
func (c *ImageController) Image() {
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
	media := models.Image{
		Id: id,
	}
	if err := media.Read("id"); err != nil {
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
