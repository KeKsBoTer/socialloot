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

	// Type of the item which is voted on (post,comment etc.)
	Type string
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

func getVotesOnPost(item string) orm.QuerySeter {
	return Votes().Filter("type", "post").Filter("item", item)
}
func getUserVoteOnPost(item string, u *User) orm.QuerySeter {
	return getVotesOnPost(item).Filter("user", u)
}
