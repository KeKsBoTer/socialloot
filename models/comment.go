package models

import (
	"time"

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
	Post    *Post    `orm:"rel(fk)"`
	ReplyTo *Comment `orm:"rel(fk);null"`
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

func (c *Comment) Valid(v *validation.Validation) {
	if c.User == nil {
		v.AddError("User", "Missing user")
	}
	if len(c.Text) < 1 {
		v.AddError("Text", "Missing comment text")
	}
	if c.Post == nil {
		v.AddError("Post", "Missing post")
	}
	if len(c.Id) != ItemIDLength {
		v.AddError("Id", "Invalid comment id")
	}
}
