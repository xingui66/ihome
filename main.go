package main

import (
	"github.com/gin-gonic/gin"
	"ihomegit/ihome/controller"
)

func main() {
	//初始化路由
	//树型寻址

	router:= gin.Default()
	//model.InitRedis()
	//model.InitDb()

	//静态路径(静态资源路径)
	router.Static("/home","view")

	//路由分组
	//动态路径(动态接口路径)
	r1:=router.Group("/api/v1.0")
	{
		r1.GET("/areas",controller.GetArea)

		r1.GET("/session",controller.GetSession)
		//传参方法,url传值,form表单传值,ajax传值,路径传值

		//传参方法,url传值,form表单传值,ajax传值,路径传值
		r1.GET("/imagecode/:uuid",controller.GetImageCd)
		r1.GET("/smscode/:mobile",controller.GetSmsCd)

		/*r1.GET("/getAllArea", func(ctx *gin.Context) {
			ctx.Writer.WriteString("喵喵")
		})*/
	}

	//model.InitModel()
    //model.InsertData();

	router.Run(":8085")
}