// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"BeegoDemo/webServer/controllers"
	"github.com/astaxie/beego"


)

type User struct {
	Id int
	username string
	password string 
	user_group string
	groupName string

}


func init() {
	beego.Router("/login/", &controllers.UserObjectController{},"post:Login")
// 	res := make(orm.Params)
// 	beego.Get("/",func(ctx *context.Context){
// 		o := orm.NewOrm()
// 		_,err := o.Raw("select username,password from main_user").RowsToMap(&res, "username", "password")
// 		if err == nil{
// 			fmt.Println(res["abcd"])
// 		}
// 		ctx.Output.Body([]byte("hello world"))
//    })

	// ns := beego.NewNamespace("/v1",
	// 	beego.NSNamespace("/object",
	// 		beego.NSInclude(
	// 			&controllers.ObjectController{},
	// 		),
	// 	),
	// 	beego.NSNamespace("/user",
	// 		beego.NSInclude(
	// 			&controllers.UserController{},
	// 		),
	// 	),
	// )
	// beego.AddNamespace(ns)
}
