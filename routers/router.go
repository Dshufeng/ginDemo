package routers

import (
	"fmt"
	 "ginDemo/controller/v1"
	"ginDemo/http/middleware"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"ginDemo/validator/member"
)

func InitRouter(r *gin.Engine) {
	v1Group := r.Group("/v1")
	{
		v1Group.GET("/ping", v1.Ping)
		v1Group.POST("/add/member", v1.AddUser)
		v1Group.GET("/sign", sign)
	}

	v2Group := r.Group("/v2").Use(middleware.SignMiddleware())
	{
		v2Group.GET("/ping", v1.Ping)
		v2Group.POST("/add/member", v1.AddUser)
	}

	if v,ok := binding.Validator.Engine().(*validator.Validate);ok{
		v.RegisterValidation("NameValid",member.NameValid)
	}
}

func social(c *gin.Context) {
	baseUrl := "http://localhost:81/api/"
	res, err := http.Get(baseUrl + "captcha")
	defer res.Body.Close()

	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(res.Body)
	c.String(200, string(body))
}


func sign(c *gin.Context) {
	ts := strconv.FormatInt(utils.GetTime(), 10)

	var params = url.Values{
		"name": []string{"li"},
		"age":  []string{"19"},
		"ts":   []string{ts},
	}

	sn := utils.CreateSign(params)
	c.JSON(200, gin.H{
		"sn": sn,
		"ts": ts,
	})
}
