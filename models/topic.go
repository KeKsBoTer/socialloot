package models

import "github.com/astaxie/beego/orm"

type Topic struct {
	Id          int    `orm:"pk;auto"`
	Name        string `orm:"unique"`
	Title       string
	Description string
	Posts       []*Post `orm:"reverse(many)"`
}

func Topics() orm.QuerySeter {
	var table Topic
	return orm.NewOrm().QueryTable(table).OrderBy("-Id")
}

func (t *Topic) Insert() error {
	if _, err := orm.NewOrm().Insert(t); err != nil {
		return err
	}
	return nil
}

func (m *Topic) Read(fields ...string) error {
	return orm.NewOrm().Read(m, fields...)
}
