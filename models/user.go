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
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}

func (u *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(u, field, fields...)
}

func (u *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) Delete() error {
	if _, err := orm.NewOrm().Delete(u); err != nil {
		return err
	}
	return nil
}

func Users() orm.QuerySeter {
	var table User
	return orm.NewOrm().QueryTable(table).OrderBy("-Id")
}

// ReadVoteOnPost gets the users vote on the given post and safes the result in the post struct
func (u *User) ReadVoteOnPost(p *Post) error {
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
func (u *User) ReadVoteOnComment(c *Comment) error {
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
