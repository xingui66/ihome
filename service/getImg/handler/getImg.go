package handler

import (
	"context"
	getImg "ihomegit/ihome/service/getImg/proto/getImg"
	"github.com/afocus/captcha"
	"image/color"
	"ihomegit/ihome/service/getImg/model"
	"ihomegit/ihome/service/getImg/utils"
	"encoding/json"
)

type GetImg struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetImg) MicroGetImg(ctx context.Context, req *getImg.Request, rsp *getImg.Response) error {
	cap:=captcha.New()
	if err := cap.SetFont("comic.ttf"); err != nil {
		panic(err.Error())
	}

	//设置验证码图片大小
	cap.SetSize(128, 64)
	//设置混淆程度
	cap.SetDisturbance(captcha.NORMAL)
	//设置字体颜色
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	//设置背景色  background
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

	//生成验证码图片
	//rand.Seed(time.Now().UnixNano())
	img,rnd := cap.Create(4,captcha.NUM)

	//存储验证码   redis
	err := model.SaveImgRnd(req.Uuid,rnd)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	imgJson , err := json.Marshal(img)
	rsp.Data= imgJson

	return nil
}

