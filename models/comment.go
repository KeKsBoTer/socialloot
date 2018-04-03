package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type Comment struct {
	Id string `orm:"pk"`

	// User is the person which wrote the comment
	User *User     `orm:"rel(fk);null;on_delete(do_nothing)"`
	Date time.Time `orm:"auto_now_add"`

	// Text is the comment
	Text string

	// Post that is commented on
	ReplyTo string

	Votes   int           `orm:"-"`
	VoteDir VoteDirection `orm:"-"`
	Replies []*Comment    `orm:"-"`
}

func Comments() orm.QuerySeter {
	var table Comment
	return orm.NewOrm().QueryTable(table)
}

func (c *Comment) Insert() error {
	if _, err := orm.NewOrm().Insert(c); err != nil {
		return err
	}
	return nil
}

func (c *Comment) Read(fields ...string) error {
	if err := orm.NewOrm().Read(c, fields...); err != nil {
		return err
	}
	return nil
}

func (c *Comment) Valid(v *validation.Validation) {
	if c.User == nil {
		v.AddError("User", "Missing user")
	}
	if len(c.Text) < 1 {
		v.AddError("Text", "Missing comment text")
	}
	if len(c.ReplyTo) != ItemIDLength {
		v.AddError("RepleyTo", "Invalid item id")
	}
	if len(c.Id) != ItemIDLength {
		v.AddError("Id", "Invalid comment id")
	}
}

func (c *Comment) ReadVoteData(u *User) error {
	if u != nil {
		// Get user vote on post
		if err := u.ReadVoteOnComment(c); err != nil {
			return err
		}
	}
	return c.ReadVoteSum()
}

// ReadVoteOnPost gets the users vote on the given post and safes the result in the post struct
func (c *Comment) ReadVoteOnPost(p *Post) error {
	var vote Vote
	if err := getVotesOnItem(c.Id).One(&vote, "action"); err != nil {
		return err
	}
	p.VoteDir = vote.Action
	return nil
}

func (c *Comment) ReadVoteSum() error {
	var votes []*Vote
	c.Votes = 0
	if _, err := getVotesOnItem(c.Id).All(&votes); err != nil {
		return err
	}
	for _, v := range votes {
		if v.Action == VoteDirectionUp {
			c.Votes++
		} else if v.Action == VoteDirectionDown {
			c.Votes--
		}
	}
	return nil
}

func (c *Comment) LoadReplies(u *User) error {
	var replies []*Comment
	if _, err := Comments().Filter("replyto", c.Id).RelatedSel("user").OrderBy("date").All(&replies); err != nil {
		return err
	}
	if err := c.ReadVoteData(u); err != nil {
		beego.Error(err)
	}
	c.Replies = replies
	// load replies recursively
	for i := range c.Replies {
		if err := c.Replies[i].LoadReplies(u); err != nil {
			beego.Error(err)
		}
	}
	return nil
}
