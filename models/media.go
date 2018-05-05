package models

import "github.com/astaxie/beego/orm"

type MediaType string

const (
	MediaImage = MediaType("image")
)

type Media struct {
	Id        int `orm:"pk;auto"`
	Type      MediaType
	File      string
	Thumbnail string
}

func (m *Media) Insert() error {
	_, err := orm.NewOrm().Insert(m)
	return err
}

func (m *Media) Read(fields ...string) error {
	return orm.NewOrm().Read(m, fields...)
}
