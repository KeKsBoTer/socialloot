package lib

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/KeKsBoTer/socialloot/models"
)

/*
 Get authenticated user and update logintime
*/
func Authenticate(name string, password string) (user *models.User, err error) {
	msg := "invalid name or password."
	user = &models.User{Name: name}

	if err := user.Read("Name"); err != nil {
		if err.Error() == "<QuerySeter> no row found" {
			err = errors.New(msg)
		}
		return user, err
	} else if user.Id < 1 {
		// No user
		return user, errors.New(msg)
	} else if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password)); err != nil {
		// No matched password
		return user, errors.New(msg)
	} else {
		user.LastLoginTime = time.Now()
		user.Update("LastLoginTime")
		return user, nil
	}
}
