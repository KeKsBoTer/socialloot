package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// User is the model for a user entry in the database
type User struct {
	Id   int    `orm:"pk;auto"`
	Name string `orm:"unique"`

	// password hashed with bcrypt
	Password string

	CreationDate  time.Time
	LastLoginTime time.Time `orm:"null"`
}

// Insert user into database
func (u *User) Insert() error {
	_, err := orm.NewOrm().Insert(u)
	return err
}

// Read user from database by given field
func (u *User) Read(fields ...string) error {
	return orm.NewOrm().Read(u, fields...)
}

// Update user data in database
// if fields are provided, only this fields are updated
func (u *User) Update(fields ...string) error {
	_, err := orm.NewOrm().Update(u, fields...)
	return err
}

// Delete user from database
func (u *User) Delete() error {
	_, err := orm.NewOrm().Delete(u)
	return err
}

// Users is a helper to query the user table
func Users() orm.QuerySeter {
	var table User
	return orm.NewOrm().QueryTable(table)
}

// ReadVoteOnPost gets the users vote on the given post and safes the result in the post struct
func (u *User) ReadVoteOnPost(p *PostMetaData) error {
	var vote Vote
	if err := u.GetVoteOnItem(p.Id).One(&vote, "action"); err != nil {
		if err != orm.ErrNoRows {
			return err
		}
	}
	p.VoteDir = vote.Action
	return nil
}

// ReadVoteOnComment gets the users vote on the given post and safes the result in the post struct
func (u *User) ReadVoteOnComment(c *CommentMetaData) error {
	var vote Vote
	if err := u.GetVoteOnItem(c.Id).One(&vote, "action"); err != nil {
		if err != orm.ErrNoRows {
			return err
		}
		vote.Action = 0
	}
	c.VoteDir = vote.Action
	return nil
}

// GetVoteOnItem is a helper to get the user's votes on a item
func (u *User) GetVoteOnItem(id string) orm.QuerySeter {
	return getVotesOnItem(id).Filter("user", u)
}
