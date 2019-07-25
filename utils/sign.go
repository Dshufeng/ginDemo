package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"ginDemo/config"
	"github.com/gin-gonic/gin"
	"net/url"
	"sort"
	"strconv"
	"time"
)

func CreateSign(params url.Values) string {
	var key []string
	var str = ""
	for k := range params {
		if k != "sn" {
			key = append(key, k)
		}
	}

	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params.Get(key[i]))
		} else {
			str += fmt.Sprintf("&%v=%v", key[i], params.Get(key[i]))
		}
	}
	//fmt.Println(str)
	sign := MD5(MD5(str) + MD5(config.APP_NAME+config.APP_SECRET))
	return sign
}

func VerifySign(c *gin.Context) {
	var method = c.Request.Method
	var req url.Values
	var sn string
	var st int64

	if method == "GET" {
		req = c.Request.URL.Query()
		sn = c.Query("sn")
		st, _ = strconv.ParseInt(c.Query("ts"), 10, 64)
	} else if method == "POST" {
		c.Request.ParseForm() // 很重要
		req = c.Request.PostForm
		sn = c.PostForm("sn")
		st, _ = strconv.ParseInt(c.PostForm("ts"), 10, 64)
	} else {
		RetErrJson(500, "request method not allowed", c)
		return
	}
	exp, _ := strconv.ParseInt(config.API_EXPIRY, 10, 64)

	if GetTime() < st || (GetTime()-st) > exp {
		RetErrJson(500, "time exp: "+strconv.FormatInt(st, 10), c)
		return
	}
	if sn == "" || sn != CreateSign(req) {
		RetErrJson(500, "invalid sn: "+sn, c)
		return
	}
}

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func GetTime() int64 {
	st := time.Now().Unix()
	return st
}

func RetErrJson(code int, msg string, c *gin.Context) {
	c.AbortWithStatusJSON(code, gin.H{
		"msg": msg,
	})
}
