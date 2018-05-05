package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id            int    `orm:"pk;auto"`
	Name          string `orm:"unique"`
	Password      string
	CreationDate  time.Time `orm:"auto_now_add`
	LastLoginTime time.Time `orm:"null"`
}

func (u *User) Insert() error {
	_, err := orm.NewOrm().Insert(u)
	return err
}

func (u *User) Read(fields ...string) error {
	return orm.NewOrm().Read(u, fields...)
}

func (u *User) Update(fields ...string) error {
	_, err := orm.NewOrm().Update(u, fields...)
	return err
}

func (u *User) Delete() error {
	_, err := orm.NewOrm().Delete(u)
	return err
}

func Users() orm.QuerySeter {
	var table User
	return orm.NewOrm().QueryTable(table).OrderBy("-Id")
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
