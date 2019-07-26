package v1

import (
	"github.com/gin-gonic/gin"
	"ginDemo/entity"
	"ginDemo/utils"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": "pong",
	})
}

func AddUser(c *gin.Context) {
	//name := c.PostForm("name")
	//age := c.DefaultPostForm("age", "11")
	var mem entity.Member
	var ret entity.Result

	if err := c.ShouldBind(&mem);err != nil{
		ret.SetCode(entity.ERR_CODE)
		ret.SetMsg(err.Error())
		c.AbortWithStatusJSON(500,ret)
		return
	}

	ret.SetCode(entity.SUCCESS_CODE)
	ret.SetMsg("success")
	ret.SetData(gin.H{
		"name": mem.Name,
		"age":  mem.Age,
	})

	c.JSON(http.StatusOK, ret)
}

func AddName(c *gin.Context)  {
	name := c.Query("name")
	var ret entity.Result

	err := validatorName(name)
	if err != nil{
		ret.Code = 400
		ret.Msg = err.Error()
		c.AbortWithStatusJSON(500,ret)
		return
	}
	ret.Code = 0
	ret.Msg = "success"
	c.JSON(200,ret)
}

func validatorName(str string)  (err error){
	if str == ""{
		err = utils.WeChat("name can not null")
	}
	return
}
