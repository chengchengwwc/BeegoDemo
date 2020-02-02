package utils

import (
	"time"
)



func GetCurrentTime() string{
	//获取当前时间
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}