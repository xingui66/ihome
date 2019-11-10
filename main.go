package main

import (
	"github.com/gin-gonic/gin"
	"ihomegit/ihome/model"
)

func main() {
	//初始化路由
	//路由分组
	router:= gin.Default()
	r0 := router.Group("/v0")
	{
		r0.GET("abc", func(ctx *gin.Context) {
			ctx.Writer.WriteString("abcdefg")
		})
	}

	r1:=router.Group("/v1")
	{
		r1.GET("edf", func(ctx *gin.Context) {
			ctx.Writer.WriteString("喵喵")
		})
	}

	model.InitModel()
    //model.InsertData();

}