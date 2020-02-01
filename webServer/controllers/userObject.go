package controllers

import (
	"BeegoDemo/webServer/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"

)

type UserObjectController struct {
	beego.Controller
}


func (u *UserObjectController) Post(){
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	fmt.Println(user)
	u.Data["json"] = map[string]string{"uid": "ddd"}
	u.ServeJSON()
}

func (u *UserObjectController) Login(){
	var user models.UserCredential
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		u.Abort("500")
		return
	}
	user_group,group_name,err := models.GetUserCredential(user.Username,user.Pwd)
	if err != nil {
		u.CustomAbort(404,"用户名不存在")
		return
	}
	fmt.Println(user_group)
	fmt.Println(group_name)
	u.Data["json"] = "登陆成功"
	u.ServeJSON()
	










	username := u.GetString("username")
	password := u.GetString("password")
	fmt.Println(username)
	fmt.Println(password)
	u.Data["json"] = "user not exist"
	u.ServeJSON()


}







