package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Topic struct {
	Id          int    `orm:"pk"`
	Name        string `orm:"unique"`
	Title       string
	Description string
	Posts       []*Post `orm:"reverse(many)"`
}

type Post struct {
	Id      string `orm:"pk"`
	User    *User  `orm:"rel(fk);null;on_delete(do_nothing)"`
	Date    time.Time
	Title   string
	Content string
	Topic   *Topic `orm:"rel(fk);null;on_delete(do_nothing)"`
}

func init() {
	orm.RegisterModel(
		new(User),
		new(Topic),
		new(Post),
	)
}
