package main

import (
	_ "BeegoDemo/webServer/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

func init(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("db1", "mysql", "root:wwc880515@tcp(127.0.0.1:3306)/mydb?charset=utf8")
}


func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
