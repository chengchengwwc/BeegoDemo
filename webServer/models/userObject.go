package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/jameskeane/bcrypt"
	"BeegoDemo/webServer/utils"
)



type UserCredential struct{
	Username string `json:"username"`
	Pwd  string `json:"password"`
	GroupName string `json:"GroupName"`
}


type LoginUserDetail struct {
	Username string `json:"username"`
	Logintime string `json:"login_time"`
}



func GetUserCredential(username string,password string) (string,string,error){
	//登陆功能
	var maps []orm.Params
	o := orm.NewOrm()
	num,err := o.Raw("SELECT password,user_group,groupName from main_user WHERE username = ?",username).Values(&maps)
	if err != nil || num == 0{
		return "","",err
	}
	if bcrypt.Match(password,maps[0]["password"].(string)){
		return maps[0]["user_group"].(string),maps[0]["groupName"].(string),nil
	}else{
		return "","",errors.New("password bad")
	}
}

func AddLoginDetail(username,session_id string) error{
	// 记录登陆时间
	login_time := utils.GetCurrentTime()
	o := orm.NewOrm()
	_,err := o.Raw("INSERT INTO sessions (login_name,login_time,token) VALUES (?,?,?)",username,login_time,session_id).Exec()
	if err != nil{
		return err
	}
	return nil
}