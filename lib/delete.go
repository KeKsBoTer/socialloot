package lib

import (
	"errors"

	"github.com/KeKsBoTer/socialloot/models"
	"github.com/astaxie/beego"
)

// DeletePost removes the given post from the database
// The user parameter is the user who wants to delete the post.
// He can only delete the post if he submited it.
func DeletePost(item string, user *models.User) error {
	post := models.Post{
		Id: item,
	}
	if err := post.Read(false); err != nil {
		beego.Error(err)
		return errors.New("Cannot read topic")
	}
	if post.User.Id != user.Id {
		return errors.New("User is not allowed to delete this post")
	}
	if err := post.Delete(); err != nil {
		beego.Error(err)
		return errors.New("An unexpected error occured")
	}
	return nil
}
