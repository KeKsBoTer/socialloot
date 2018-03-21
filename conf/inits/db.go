package inits

import (
	_ "github.com/KeKsBoTer/socialloot/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/mattn/go-sqlite3"
)

func init() {

	dbname := "default"
	runmode := beego.AppConfig.String("runmode")
	datasource := beego.AppConfig.String("datasource")

	switch runmode {
	case "dev":
		orm.Debug = true
		fallthrough
	default:
		orm.RegisterDriver("sqlite3", orm.DRSqlite)
		orm.RegisterDataBase(dbname, "sqlite3", datasource)
	}

	force, verbose := false, true
	err := orm.RunSyncdb(dbname, force, verbose)
	if err != nil {
		panic(err)
	}

	// orm.RunCommand()
}
