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
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Media) Read(fields ...string) error {
	return orm.NewOrm().Read(m, fields...)
}
