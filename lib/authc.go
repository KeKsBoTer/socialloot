package lib

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/KeKsBoTer/socialloot/models"
)

// Authenticate checks if the given user credentials are valid
// If the user name is found in the database the password hashes are compared
// Also the LastLoginTime time field for the user is updates
func Authenticate(name string, password string) (user *models.User, err error) {
	msg := "Invalid name or password."
	user = &models.User{Name: name}
	if err := user.Read("Name"); err != nil {
		if err.Error() == "<QuerySeter> no row found" {
			err = errors.New(msg)
		}
		return nil, err
	} else if user.Id < 1 {
		// No user
		return nil, errors.New(msg)
	} else if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// No matched password
		return nil, errors.New(msg)
	}
	user.LastLoginTime = time.Now()
	user.Update("LastLoginTime")
	return user, nil
}
