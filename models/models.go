package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(
		new(User),
		new(Topic),
		new(Post),
		new(Vote),
		new(Comment),
		new(Media),
	)
}
