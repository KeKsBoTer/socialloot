package lib

import (
	"errors"
	"time"

	"github.com/astaxie/beego"

	"github.com/KeKsBoTer/socialloot/models"
	"golang.org/x/crypto/bcrypt"
)

func SignupUser(username, password string) (*models.User, error) {
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		beego.Error(err)
		return nil, errors.New("Something seems to be wrong with your password")
	}
	u := models.User{
		Name:          username,
		Password:      string(hashedPw),
		LastLoginTime: time.Now(),
	}
	if err := u.Insert(); err != nil {
		return nil, err
	}
	return &u, nil
}
