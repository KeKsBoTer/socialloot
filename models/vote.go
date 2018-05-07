package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

// VoteDirection is the direction of a vote
type VoteDirection int

const (
	// VoteDirectionUp means upvote
	VoteDirectionUp = VoteDirection(1)

	// VoteDirectionDown means downvote
	VoteDirectionDown = VoteDirection(-1)
)

// String converts direction to "upvote" or "downvote"
func (v *VoteDirection) String() string {
	switch *v {
	case VoteDirectionUp:
		return "upvote"
	case VoteDirectionDown:
		return "downvote"
	default:
		return strconv.Itoa(int(*v))
	}
}

// Vote is the model for a vote in the database
type Vote struct {
	Id int `orm:"pk;auto"`

	// The user that executed the vote
	User *User `orm:"rel(fk);null;on_delete(do_nothing)"`

	Date time.Time `orm:"auto_now"`

	// Action performed by User (down- or upvote)
	Action VoteDirection

	// Item that is voted on
	Item string `orm:"size(11)"`
}

// InsertOrUpdate adds vote to database or updates the vote's direction
// if the user allready voted on this item
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
	if _, err := o.Update(v, "action"); err != nil {
		return err
	}
	return nil
}

// Votes is helper to query votes
func Votes() orm.QuerySeter {
	var table Vote
	return orm.NewOrm().QueryTable(table)
}

// helper to query all votes on a item
func getVotesOnItem(id string) orm.QuerySeter {
	return Votes().Filter("item", id)
}
