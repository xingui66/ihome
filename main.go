package main

import (
	"github.com/gin-gonic/gin"
	"ihomegit/ihome/controller"
	"ihomegit/ihome/model"
	"fmt"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-contrib/sessions"
	"ihomegit/ihome/utils"
	"net/http"
	"time"
)

func main() {
	//初始化路由
	//树型寻址

	router:= gin.Default()
	model.InitRedis()
	err := model.InitDb()
	if  err != nil {
		//把错误打印到日志文件中
		fmt.Println(err)
		return
	}

	//初始化redis容器,存储session数据
	store,_ :=redis.NewStore(20,"tcp","127.0.0.1:6379","",[]byte("session"))
	//静态路径(静态资源路径)
	router.Static("/home","view")

	/*router.GET("/testCookie", func(context *gin.Context) {
		context.SetCookie("testCookie","喵喵",0,"路径?","domain?",false,true)
		context.Writer.WriteString("测试下cookie")
	})*/

	//路由分组
	//动态路径(动态接口路径)
	r1:=router.Group("/api/v1.0")
	{
		r1.GET("/areas",controller.GetArea)

		//r1.GET("/session",controller.GetSession)
		//传参方法,url传值,form表单传值,ajax传值,路径传值

		//传参方法,url传值,form表单传值,ajax传值,路径传值
		r1.GET("/imagecode/:uuid",controller.GetImageCd)
		r1.GET("/smscode/:mobile",controller.GetSmsCd)
		r1.POST("/users",controller.PostRet)//注册
		/*r1.GET("/getAllArea", func(ctx *gin.Context) {
			ctx.Writer.WriteString("喵喵")
		})*/

		//登录业务
		//r1.Use(sessions.Sessions("mysession",store))
		//登录业务   路由过滤器   中间件
		r1.Use(sessions.Sessions("mysession",store))
		r1.POST("/sessions",controller.PostLogin)//登录
		r1.GET("/session",controller.GetSession)

		r1.Use(Filter)

		r1.DELETE("/session",controller.DeleteSession)
		r1.GET("/user",controller.GetUserInfo)
		r1.PUT("/user",controller.PutUserInfo)

	}



	//model.InitModel()
    //model.InsertData();

	router.Run(":8085")
}


//路由过滤器


func Filter(ctx *gin.Context){

	session :=sessions.Default(ctx)
	userName := session.Get("userName")
	resp:= make(map[string]interface{})
	if userName ==nil {
		resp["errno"]= utils.RECODE_SESSIONERR
		resp["errmsg"]= utils.RecodeText(utils.RECODE_SESSIONERR)
		ctx.JSON(http.StatusOK,resp)
		ctx.Abort()
		return
	}

	fmt.Println("next之前打印",time.Now())
	ctx.Next()
	fmt.Println("next之后打印...",time.Now())


}