package lib

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/KeKsBoTer/socialloot/models"
)

func SignupUser(u *models.User) (int, error) {
	var (
		err error
		msg string
	)

	if models.Users().Filter("name", u.Email).Exist() {
		msg = "was already regsitered input name address."
		return 0, errors.New(msg)
	}

	hashedPw, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	u.Password = string(hashedPw)
	u.CreationDate = time.Now()
	u.LastLoginTime = u.CreationDate
	err = u.Insert()
	if err != nil {
		return 0, err
	}

	return u.Id, err
}
