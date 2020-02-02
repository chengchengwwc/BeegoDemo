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


func CheckUserCredential(username string) (int64,error){
	// 检查用户名是否存在
	o := orm.NewOrm()
	res, err := o.Raw("SELECT username FROM main_user WHERE username = ?",username).Exec()
	if err == nil{
		num,_ := res.RowsAffected()
		return num, nil
	}else{
		return 0,err
	}
}


func AddUseCredential(username,password,group,team string)(string,error) {
	num,err := CheckUserCredential(username)
	if err != nil{
		return "",err
	}
	if num > 0{
		return "用户名已经存在",nil
	}
	hash, err := bcrypt.Hash(password)
	if err != nil{
		return "密码哈希化失败",err
	}
	o := orm.NewOrm()
	_,err = o.Raw("INSERT INTO main_user (username,password,user_group,groupName) VALUES (?,?,?,?)",username,hash,group,team).Exec()
	if err != nil{
		return "用户创建失败",err
	}
	return "用户创建成功",nil
}