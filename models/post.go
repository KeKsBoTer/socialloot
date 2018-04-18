package models

import (
	"time"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

// The length of all item ids
// Change in model tags!!
const ItemIDLength = 11

type PostType string

const (
	PostTypeImage = PostType("image")
	PostTypeText  = PostType("text")
	PostTypeLink  = PostType("link")
)

type Post struct {
	Id    string    `orm:"pk;size(11)"`
	User  *User     `orm:"rel(fk);null;on_delete(do_nothing)"`
	Date  time.Time `orm:"auto_now_add"`
	Title string

	// Content depends on the type of the post.
	// Text and Link is stored as plain text in the content field
	// For images the image id (see table media) is saved in the field
	Content string

	// Type of the post. Image, text, or link
	Type PostType

	Topic *Topic `orm:"rel(fk);null;on_delete(do_nothing)"`

	// for html rendering
	Votes    int           `orm:"-"`
	VoteDir  VoteDirection `orm:"-"`
	Comments []*Comment    `orm:"-"`
}

func Posts() orm.QuerySeter {
	var table Post
	return orm.NewOrm().QueryTable(table)
}

func (p *Post) Insert() error {
	if _, err := orm.NewOrm().Insert(p); err != nil {
		return err
	}
	return nil
}

func (p *Post) Read(fields ...string) error {
	return orm.NewOrm().Read(p, fields...)
}

func (p *Post) ReadVoteData(u *User) error {
	if u != nil {
		// Get user vote on post
		if err := u.ReadVoteOnPost(p); err != nil {
			return err
		}
	}
	return p.ReadVoteSum()
}

func (p *Post) ReadVoteSum() error {
	var votes []*Vote
	p.Votes = 0
	if _, err := getVotesOnItem(p.Id).All(&votes); err != nil {
		return err
	}
	for _, v := range votes {
		if v.Action == VoteDirectionUp {
			p.Votes++
		} else if v.Action == VoteDirectionDown {
			p.Votes--
		}
	}
	return nil
}

func (p *Post) ReadComments(u *User) error {
	if _, err := Comments().Filter("replyto", p.Id).RelatedSel("user").All(&p.Comments); err != nil {
		return err
	}
	for _, c := range p.Comments {
		if err := c.ReadVoteData(u); err != nil {
			beego.Error(err)
		}
		if err := c.LoadReplies(u); err != nil {
			beego.Error(err)
		}
	}
	return nil
}
