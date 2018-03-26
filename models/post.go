package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type Post struct {
	Id        string `orm:"pk"`
	User      *User  `orm:"rel(fk);null;on_delete(do_nothing)"`
	Date      time.Time
	Title     string
	Content   string `form:"Content" valid:"Required"`
	TopicName string `orm:"-" form:"Topic" valid:"Required"`
	Topic     *Topic `orm:"rel(fk);null;on_delete(do_nothing)"`
	Votes int `orm:"-"`
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

func (p *Post) Valid(v *validation.Validation) {
	if p.User == nil {
		v.AddError("User", "Missing user")
	}
	if len(p.Title) < 1 {
		v.AddError("Title", "Missing title")
	}
	if len(p.Content) < 1 {
		v.AddError("Content", "Missing content")
	}
	if len(p.TopicName) < 1 {
		v.AddError("Topic", "Missing topic")
	}
}

func (p *Post) Read(fields ...string) error {
	if err := orm.NewOrm().Read(p, fields...); err != nil {
		return err
	}
	return nil
}
