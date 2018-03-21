package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type User struct {
	Id            int    `orm:"pk;auto"`
	Name          string `orm:"unique" form:"Name" valid:"Required;Name"`
	RePassword    string `orm:"-" form:"Repassword" valid:"Required"`
	Password      string `orm:"size(32)" form:"Password" valid:"Required;MinSize(8)"`
	CreationDate  time.Time
	LastLoginTime time.Time `orm:"null"`
	Email         string    `orm:"null"`
}

func (u *User) Valid(v *validation.Validation) {
	if u.Password != u.RePassword {
		v.SetError("Repassword", "Does not matched password, repassword")
	}
}

func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func Users() orm.QuerySeter {
	var table User
	return orm.NewOrm().QueryTable(table).OrderBy("-Id")
}
