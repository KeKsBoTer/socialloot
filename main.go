package main

import (
	_ "github.com/KeKsBoTer/socialloot/conf/inits"
	_ "github.com/KeKsBoTer/socialloot/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
