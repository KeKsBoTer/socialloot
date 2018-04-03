package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type VoteDirection int

const (
	VoteDirectionUp   = VoteDirection(1)
	VoteDirectionDown = VoteDirection(-1)
)

type Vote struct {
	Id   int       `orm:"pk;auto"`
	User *User     `orm:"rel(fk);null;on_delete(do_nothing)"`
	Date time.Time `orm:"auto_now"`

	// Action performed by User (down- or upvote)
	Action VoteDirection

	// Item that is voted on
	Item string `orm:"size(11)"`
}

func (v *Vote) InsertOrUpdate() error {
	o := orm.NewOrm()
	newAction := v.Action
	if err := o.Read(v, "user", "item"); err != nil {
		if err == orm.ErrNoRows {
			_, err = o.Insert(v)
			return err
		}
		return err
	}
	v.Action = newAction
	if _, err := o.Update(v); err != nil {
		return err
	}
	return nil
}

func Votes() orm.QuerySeter {
	var table Vote
	return orm.NewOrm().QueryTable(table)
}

func getVotesOnItem(id string) orm.QuerySeter {
	return Votes().Filter("item", id)
}

func (u *User) GetVoteOnItem(id string) orm.QuerySeter {
	return getVotesOnItem(id).Filter("user", u)
}
