package handler

import (
	"context"


	getSms "ihomegit/ihome/service/getSms/proto/getSms"
	"ihomegit/ihome/service/getSms/model"
	"ihomegit/ihome/utils"
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"math/rand"
	"fmt"
	"time"
)

type GetSms struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetSms) MicroGetSms(ctx context.Context, req *getSms.Request, rsp *getSms.Response) error {

	//获取验证码,判断验证码是否正确,如果不正确,就不发送短信
	rnd,err := model.GetImgCode(req.Uuid)
	if err != nil {
		rsp.Errno = utils.RECODE_NODATA
		rsp.Errmsg = utils.RecodeText(utils.RECODE_NODATA)
		return err
	}

	if req.Text !=rnd{
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		//返回自定义的error数据
		return errors.New("验证码输入错误")
	}
	//如果成功,发送短信,存储短信验证码  阿里云短信接口
	client, err := sdk.NewClientWithAccessKey("default", "LTAI4FexwrAFbn4ua4DHAyXh", "AltI2inQ1I5TqAEwAfrJNgP54VnVOx")
	if err != nil {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return err
	}
	//获取6位数随机码
	myRnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06d", myRnd.Int31n(1000000))

	//初始化请求对象
	request := requests.NewCommonRequest()
	request.Method = "POST"//设置请求方法
	request.Scheme = "https" // https | http   //设置请求协议
	request.Domain = "dysmsapi.aliyuncs.com"  //域名
	request.Version = "2017-05-25"			//版本号
	request.ApiName = "SendSms"				//api名称
	request.QueryParams["PhoneNumbers"] = req.Mobile  //需要发送的电话号码
	request.QueryParams["SignName"] = "北京5期区块链"    //签名名称   需要申请
	request.QueryParams["TemplateCode"] = "SMS_176375357"   //模板号   需要申请
	request.QueryParams["TemplateParam"] = `{"code":`+vcode+`}`   //发送短信验证码

	response, err := client.ProcessCommonRequest(request)  //发送短信
	//如果不成功
	if !response.IsSuccess(){
		rsp.Errno = utils.RECODE_SMSERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_SMSERR)
		return errors.New("发送短信失败")
	}

	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	return nil
}
