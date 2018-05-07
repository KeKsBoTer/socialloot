package lib

import (
	"errors"
	"time"

	"github.com/astaxie/beego"

	"github.com/KeKsBoTer/socialloot/models"
	"golang.org/x/crypto/bcrypt"
)

// SignupUser creates user in database and returns the full user data
// The password is hashed with bcrypt
func SignupUser(username, password string) (*models.User, error) {
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		beego.Error(err)
		return nil, errors.New("Something seems to be wrong with your password")
	}
	now := time.Now()
	u := models.User{
		Name:          username,
		Password:      string(hashedPw),
		LastLoginTime: now,
		CreationDate:  now,
	}
	if err := u.Insert(); err != nil {
		return nil, err
	}
	return &u, nil
}
