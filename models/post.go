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
	Content   string
	TopicName string `orm:"-" form:"Topic" valid:"Required"`
	Topic     *Topic `orm:"rel(fk);null;on_delete(do_nothing)"`
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
		v.SetError("User", "Missing user")
	}
}
