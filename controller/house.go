package controller

import (
	"github.com/gin-gonic/gin"
	getArea "ihomegit/ihome/proto/getArea"
	getImg "ihomegit/ihome/proto/getImg"
	"github.com/micro/go-micro/client"
	"context"
	"fmt"
	"net/http"
	"github.com/afocus/captcha"
	"encoding/json"
	"image/png"
	"ihomegit/ihome/utils"
	getSms "ihomegit/ihome/proto/getSms"
//	"github.com/hashicorp/consul/command/services/register"
)

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
	microClient:=getArea.NewGetAreaService("go.micro.srv.getArea",client.DefaultClient)
    resp,err :=microClient.MicroGetArea(context.TODO(),&getArea.Request{})
	if err != nil {
		fmt.Println("microCLient.MicroGetArea() err :",err)
	}
	ctx.JSON(http.StatusOK,resp)

}

//写一个假的session请求返回
func GetSession(ctx *gin.Context) {
	//构造未登录
	resp := make(map[string]interface{})

	resp["errno"] = utils.RECODE_LOGINERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_LOGINERR)

	ctx.JSON(http.StatusOK,resp)

}

func GetImageCd(ctx *gin.Context){

	uuid :=ctx.Param("uuid")
	if uuid ==""{
		fmt.Println("获取数据错误!")
		return
	}

	microClient:=getImg.NewGetImgService("go.micro.srv.getImg",client.DefaultClient)
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
		fmt.Println("调用远程服务错误!")
		return
	}
	ctx.JSON(http.StatusOK,resp)

}
