package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// The length of all item ids
// Change in model tags!!
const ItemIDLength = 11

type Post struct {
	Id      string    `orm:"pk;size(11)"`
	User    *User     `orm:"rel(fk);null;on_delete(do_nothing)"`
	Date    time.Time `orm:"auto_now_add"`
	Title   string
	Content string
	Topic   *Topic `orm:"rel(fk);null;on_delete(do_nothing)"`

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
	if err := orm.NewOrm().Read(p, fields...); err != nil {
		return err
	}
	return nil
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

// ReadVoteOnPost gets the users vote on the given post and safes the result in the post struct
func (u *User) ReadVoteOnPost(p *Post) error {
	var vote Vote
	if err := getUserVoteOnPost(p.Id, u).One(&vote, "action"); err != nil {
		return err
	}
	p.VoteDir = vote.Action
	return nil
}

func (p *Post) ReadVoteSum() error {
	var votes []*Vote
	p.Votes = 0
	if _, err := getVotesOnPost(p.Id).All(&votes); err != nil {
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

func (p *Post) ReadComments() error {
	if _, err := Comments().Filter("post", p).RelatedSel("user").All(&p.Comments); err != nil {
		return err
	}
	return nil
}
