package inits

import (
	"time"

	_ "github.com/KeKsBoTer/socialloot/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	dbname := "default"
	datasource := beego.AppConfig.String("datasource") // sqlite file

	orm.DefaultTimeLoc = time.Local
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase(dbname, "sqlite3", datasource)

	// sync model and database
	if err := orm.RunSyncdb(dbname, false, true); err != nil {
		panic(err)
	}
}
