package controllers

import (
	"github.com/astaxie/beego/logs"
	"BeegoDemo/webServer/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
	"BeegoDemo/webServer/utils"

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

// @Title Login
// @Description  用户登陆
// @Param	body username password
// @Success 200 {int} token，username，group,groupName
// @Failure 404 用户名不存在
// @Failure 500 数据序列化失败
// @router /login/ [post]
func (u *UserObjectController) Login(){
	var user models.UserCredential
	messageMap := make(map[string]string)
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		logs.Error("用户登陆数据序列化失败：",err)
		u.Abort("500")
		return
	}
	user_group,group_name,err := models.GetUserCredential(user.Username,user.Pwd)
	if err != nil {
		logs.Error("用户名不存在",err)
		u.CustomAbort(404,"用户名不存在")
		return
	}
	token := utils.GetSessionToken()
	models.AddLoginDetail(user.Username,utils.GetSessionToken())
	messageMap["token"] = token
	messageMap["username"] = user.Username
	messageMap["group"] = user_group
	messageMap["GroupName"] = group_name
	u.Data["json"] = messageMap
	u.ServeJSON()
}







