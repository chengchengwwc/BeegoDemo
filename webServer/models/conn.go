package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)


func init(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("default", "mysql", "root:wwc880515@tcp(127.0.0.1:3306)/mydb?charset=utf8")
}