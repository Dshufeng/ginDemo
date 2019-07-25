package v1

import (
	"github.com/gin-gonic/gin"
	"ginDemo/entity"
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

	c.JSON(200, ret)
}
