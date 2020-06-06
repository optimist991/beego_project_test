package main

import (
	_ "beego_project_test/routers"
	services "beego_project_test/services/db"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	services.InitSql()
	beego.Run()
}
