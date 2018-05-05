package models

import "github.com/astaxie/beego/orm"

type Topic struct {
	Id          int    `orm:"pk;auto"`
	Name        string `orm:"unique"`
	Title       string
	Description string
	Posts       []*Post `orm:"reverse(many)"`
}

func ReadTopic(name string) (*Topic, error) {
	t := Topic{Name: name}
	return &t, t.Read("name")
}

func Topics() orm.QuerySeter {
	var table Topic
	return orm.NewOrm().QueryTable(table).OrderBy("-Id")
}

func (t *Topic) Insert() error {
	_, err := orm.NewOrm().Insert(t)
	return err
}

func (m *Topic) Read(fields ...string) error {
	return orm.NewOrm().Read(m, fields...)
}
