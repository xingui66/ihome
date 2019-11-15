package controller

import (
	"github.com/gin-gonic/gin"
	getArea "ihomegit/ihome/proto/getArea"
	getImg "ihomegit/ihome/proto/getImg"
	"context"
	"fmt"
	"net/http"
	"github.com/afocus/captcha"
	"encoding/json"
	"image/png"
	"ihomegit/ihome/utils"
	getSms "ihomegit/ihome/proto/getSms"
	registerandLogin "ihomegit/ihome/proto/RegAndLog"
	updateUser "ihomegit/ihome/service/updateUser/proto/updateUser"
//	"github.com/hashicorp/consul/command/services/register"
	//"github.com/hashicorp/consul/command/services/register"
	"github.com/gin-contrib/sessions"

)

type RegStu struct {
	Mobile string `json:"mobile"`
	PassWord string `json:"password"`
	SmsCode string  `json:"sms_code"`
}


func GetArea(ctx *gin.Context){
	/*resp:=make(map[string] interface{})
	defer ctx.JSON(http.StatusOK,resp)

	areas,err := model.QueryArea();
	if err != nil {
		fmt.Println("model.GetArea err:")
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"]=utils.RecodeText(utils.RECODE_DBERR)
		return
	}

	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = areas*/
	microClient:=getArea.NewGetAreaService("go.micro.srv.getArea",utils.GetMicroClient())
    resp,err :=microClient.MicroGetArea(context.TODO(),&getArea.Request{})
	if err != nil {
		fmt.Println("microCLient.MicroGetArea() err :",err)
	}
	ctx.JSON(http.StatusOK,resp)

}

//写一个假的session请求返回
func GetSession(ctx *gin.Context) {
	/*//构造未登录
	resp := make(map[string]interface{})

	resp["errno"] = utils.RECODE_LOGINERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_LOGINERR)

	ctx.JSON(http.StatusOK,resp)*/

	resp := make(map[string]interface{})
	session:=sessions.Default(ctx)
	userName:=session.Get("userName")
	if userName == nil{
		resp["errno"] = utils.RECODE_LOGINERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_LOGINERR)
	}else {
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

		//可以是结构体,可以是map
		tempMap := make(map[string]interface{})
		tempMap["name"] = userName.(string)
		resp["data"] = tempMap
	}

	ctx.JSON(http.StatusOK,resp)

}

func GetImageCd(ctx *gin.Context){

	uuid :=ctx.Param("uuid")
	if uuid ==""{
		fmt.Println("获取数据错误!")
		return
	}

	microClient:=getImg.NewGetImgService("go.micro.srv.getImg",utils.GetMicroClient())
	resp,err :=microClient.MicroGetImg(context.TODO(),&getImg.Request{Uuid:uuid})
	if err != nil {
		fmt.Println("microCLient.MicroGetImg() err :",err)
		fmt.Println("获取远端数据失败")
		ctx.JSON(http.StatusOK,resp)
		return
	}

	var img captcha.Image
	json.Unmarshal(resp.Data,&img)
	png.Encode(ctx.Writer,img)


}

func GetSmsCd(ctx *gin.Context){
	mobile := ctx.Param("mobile")
	text:=ctx.Query("text")
	uuid :=ctx.Query("id")

	//校验数据
	if mobile == "" || text == "" || uuid ==""{
		fmt.Println("传入数据不完整")
		return
	}

	microClient :=getSms.NewGetSmsService("go.micro.srv.getSms",utils.GetMicroClient())

	resp,err:=microClient.MicroGetSms(context.TODO(),&getSms.Request{
		Uuid:uuid,
		Text:text,
		Mobile:mobile,
	})

	if err!=nil{
		fmt.Println("短信发送失败!resp:",resp,"err:",err)
		fmt.Println("调用远程服务错误!")
		return
	}
	ctx.JSON(http.StatusOK,resp)

}

//注册
func PostRet(ctx *gin.Context){
	var reg RegStu
	err :=ctx.Bind(&reg)
	if err != nil {
		fmt.Println("获取前端传递过来的数据失败")
		return
	}

	microClient:=registerandLogin.NewRegAndLogService("go.micro.srv.RegAndLog",utils.GetMicroClient())
    resp,err:=microClient.MiroRegister(context.TODO(),&registerandLogin.Request{
    	Mobile:reg.Mobile,
    	Password:reg.PassWord,
    	Smscode:reg.SmsCode,
	})

	if err != nil {
		fmt.Println("house.go microClient.MiroRegister 调用远程服务错误 err:",err)
	}
	ctx.JSON(http.StatusOK,resp)
}


type LogStu struct {
	Mobile string `json:"mobile"`
	PassWord string `json:"password"`
}

//登录
func PostLogin(ctx *gin.Context){
	//获取数据
	var log LogStu
	err := ctx.Bind(&log)
	//校验数据
	if err != nil {
		fmt.Println("获取数据失败")
		return
	}
	//处理数据   把业务放在为服务中
	//初始化客户端

	microClient := registerandLogin.NewRegAndLogService("go.micro.srv.RegAndLog",utils.GetMicroClient())
	//调用远程服务
	resp,err := microClient.MiroLogin(context.TODO(),&registerandLogin.Request{
		Mobile:log.Mobile,
		Password:log.PassWord,
	})
	defer ctx.JSON(http.StatusOK,resp)
	if err != nil {
		fmt.Println("调用login服务错误",err)
		return
	}


	//返回数据  存储session  并返回数据给web端
	session := sessions.Default(ctx)
	session.Set("userName",resp.Name)
	session.Save()

}

//删除登录记录session
func DeleteSession(ctx * gin.Context){
	session:=sessions.Default(ctx)
	resp := make(map[string]interface{})
	fmt.Println("控制器函数执行....")
	session.Delete("userName")
	err:=session.Save()
	defer ctx.JSON(http.StatusOK,resp)
	if err != nil {
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
		return
	}

	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

}


type GetUserParam struct {
}
func GetUserInfo(ctx * gin.Context){
	session := sessions.Default(ctx)
	userName := session.Get("userName")

	updataUserClient :=updateUser.NewUpdateUserService("go.micro.srv.updateUser",utils.GetMicroClient())
	resp,err:=updataUserClient.MicroGetUserInfo(context.TODO(),&updateUser.Request{
		Name:userName.(string),
	})

	if err != nil {
		fmt.Println("调用远程user服务错误",err)
		resp.Errno = utils.RECODE_DATAERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
	}

	ctx.JSON(http.StatusOK,resp)
}

type UpdateStu struct {
	Name string `json:name`
}

func PutUserInfo (ctx *gin.Context){
	var nameData UpdateStu
	err := ctx.Bind(&nameData)
	if err != nil {
		fmt.Println("获取修改的name失败")
		return
	}

    session := sessions.Default(ctx)
    userName := session.Get("userName")
    microClient := updateUser.NewUpdateUserService("go.micro.srv.updateUser",utils.GetMicroClient())
    resp,_ :=microClient.MicroUpdateUser(context.TODO(),&updateUser.UpdateReq{
    	NewName:nameData.Name,
    	OldName:userName.(string),
	})

	//更新session数据
	if resp.Errno == utils.RECODE_OK{
		//更新成功,session中的用户名也需要更新一下
		session.Set("userName",nameData.Name)
		session.Save()
	}

	ctx.JSON(http.StatusOK,resp)


}

