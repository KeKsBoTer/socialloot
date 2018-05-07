package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	// register models to orm
	orm.RegisterModel(
		new(User),
		new(Topic),
		new(Post),
		new(Vote),
		new(Comment),
		new(Image),
	)
}
