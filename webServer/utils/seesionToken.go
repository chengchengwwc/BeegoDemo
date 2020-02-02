package utils

import (
	"crypto/md5"
	"time"
	"io"
	"strconv"
	"fmt"
)


func GetSessionToken() string {
	//根据时间戳生产session token
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}



