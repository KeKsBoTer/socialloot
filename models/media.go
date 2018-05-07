package models

import "github.com/astaxie/beego/orm"

// Image is the model for a image entry in the database
// The images are stored as BLOB in the file and thumbnail field.
type Image struct {

	// Unique id for image
	Id int `orm:"pk;auto"`

	// The image as string
	File string

	// Small version of image as string
	Thumbnail string
}

// Insert image to database
func (m *Image) Insert() error {
	_, err := orm.NewOrm().Insert(m)
	return err
}

// Read image from database by given field
func (m *Image) Read(fields ...string) error {
	return orm.NewOrm().Read(m, fields...)
}
