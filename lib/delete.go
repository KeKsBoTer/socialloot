package lib

import (
	"errors"

	"github.com/KeKsBoTer/socialloot/models"
	"github.com/astaxie/beego"
)

func DeletePost(item string, user *models.User) error {
	post := models.Post{
		Id: item,
	}
	if err := post.Read(); err != nil {
		beego.Error(err)
		return errors.New("Cannot read topic")
	}
	if post.User.Id != user.Id {
		return errors.New("User is not allowed to delete post")
	}
	if err := post.Delete(); err != nil {
		beego.Error(err)
		return errors.New("An unexpected error occured")
	}
	return nil
}
