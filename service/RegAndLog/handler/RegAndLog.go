package handler

import (
	"context"
	regAndLog "ihomegit/ihome/service/RegAndLog/proto/RegAndLog"
	"ihomegit/ihome/service/RegAndLog/model"
	"ihomegit/ihome/service/RegAndLog/utils"
	"errors"
	"crypto/sha256"
	"fmt"
)

type Register struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Register) MiroRegister(ctx context.Context, req *regAndLog.Request, rsp *regAndLog.Response) error {
	//校验短信验证码会否正确
	smsCode,err := model.GetSmsCode(req.Mobile)
	if err != nil {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return err
	}

	if smsCode != req.Smscode{
		rsp.Errno = utils.RECODE_SMSERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_SMSERR)
		return errors.New("验证码错误")
	}

	pwdByte := sha256.Sum256([]byte(req.Password))
	pwd_hash := string(pwdByte[:])
	pwdHash := fmt.Sprintf("%x",pwd_hash)

	//保存到数据库
	err = model.SaveUser(req.Mobile,pwdHash)
	if err != nil {
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	return nil
}
func (e*Register)MiroLogin(ctx context.Context, req *regAndLog.Request, rsp *regAndLog.Response) error {
	//查询输入手机号和密码是否正确  mysql
	//给密码加密
	pwdByte := sha256.Sum256([]byte(req.Password))
	pwd_hash := string(pwdByte[:])
	//要把sha256得到的数据转换之后存储  转换16进制的
	pwdHash := fmt.Sprintf("%x",pwd_hash)

	user,err := model.CheckUser(req.Mobile,pwdHash)
	if err != nil {
		rsp.Errno = utils.RECODE_LOGINERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_LOGINERR)
		return err
	}

	//查询成功  登录成功  把用户名存储到session中  把用户名传给web端
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	rsp.Name = user.Name

	return nil
}