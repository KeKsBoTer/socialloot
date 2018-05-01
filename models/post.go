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
}

// ReadPost reads given post from database
func ReadPost(id string, loadRelated bool) (*Post, error) {
	p := Post{Id: id}
	return &p, p.Read(loadRelated, "id")
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

func (p *Post) Delete() error {
	if _, err := orm.NewOrm().Delete(p); err != nil {
		return err
	}
	return nil
}

func (p *Post) Read(loadRelated bool, fields ...string) error {
	o := orm.NewOrm()
	if err := o.Read(p, fields...); err != nil {
		return err
	}
	// read user and topic data for post
	if loadRelated {
		if _, err := o.LoadRelated(p, "user"); err != nil {
			return err
		}
		if _, err := o.LoadRelated(p, "topic"); err != nil {
			return err
		}
	}
	return nil
}

// PostList is a list of posts
type PostList []*Post

// PostMetaDataList is a list of post meta data
type PostMetaDataList []*PostMetaData

// ToMetaData turns a list of posts into a list of  PostMetaData
// Important: None of the meta data is loaded!
func (p *PostList) ToMetaData() *PostMetaDataList {
	metas := make(PostMetaDataList, len(*p))
	for i, post := range *p {
		metas[i] = post.NewMetaData()
	}
	return &metas
}

func (p *PostMetaDataList) ReadVoteData(u *User) error {
	for _, i := range *p {
		if err := i.ReadVoteData(u); err != nil {
			return err
		}
	}
	return nil
}

// PostMetaData is a wrapper for the post model and adds additional fields like votes and a list of comments to the model.
// All added fields are not stored in the database and are calculated at runtime
type PostMetaData struct {
	*Post
	Votes    int
	Rank     float64
	VoteDir  VoteDirection
	Comments []*CommentMetaData
}

// NewMetaData creates a new MetaData wrapper from a post
// This sould be the only way new meta data objects are created!
func (p *Post) NewMetaData() *PostMetaData {
	return &PostMetaData{
		Post: p,
	}
}

func (p *PostMetaData) ReadVoteData(u *User) error {
	if u != nil {
		// Get user vote on post
		if err := u.ReadVoteOnPost(p); err != nil {
			return err
		}
	}
	return p.ReadVoteSum()
}

func (p *PostMetaData) ReadVoteSum() error {
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
		} else {
			beego.Error("Invalid vote direction for vote", v.Id, ":", v.Action)
		}
	}
	return nil
}

func (p *PostMetaData) ReadComments(u *User) error {
	var comments CommentList
	if _, err := Comments().Filter("replyto", p.Id).RelatedSel("user").All(&comments); err != nil {
		return err
	}
	p.Comments = *comments.ToMetaData()
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
