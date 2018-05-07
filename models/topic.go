package models

import "github.com/astaxie/beego/orm"

// Topic is the model for a topic in the database
type Topic struct {

	// The topic's id (machine key)
	Id int `orm:"pk;auto"`

	// unique short name for the topic
	Name string `orm:"unique"`

	Title       string
	Description string

	// Reference to all posts submitted to this topic
	Posts []*Post `orm:"reverse(many)"`
}

// ReadTopic reads topic from database by name
func ReadTopic(name string) (*Topic, error) {
	t := Topic{Name: name}
	return &t, t.Read("name")
}

// Topics is a helper to query the topics table
func Topics() orm.QuerySeter {
	var table Topic
	return orm.NewOrm().QueryTable(table).OrderBy("-Id")
}

// Insert a topic to database
func (t *Topic) Insert() error {
	_, err := orm.NewOrm().Insert(t)
	return err
}

// Read topic from database by the given field
func (t *Topic) Read(fields ...string) error {
	return orm.NewOrm().Read(t, fields...)
}
