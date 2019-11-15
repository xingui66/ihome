package handler

import (
	"context"

	updateUser "ihomegit/ihome/service/updateUser/proto/updateUser"
	"ihomegit/ihome/service/updateUser/model"
	"ihomegit/ihome/service/updateUser/utils"
)

type GetUserInfo struct{}


func (e *GetUserInfo) MicroGetUserInfo(ctx context.Context, req *updateUser.Request, rsp *updateUser.Response) error {

	myUser ,err := model.GetUserInfo(req.Name)
	if err != nil {
		rsp.Errno = utils.RECODE_USERERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_USERERR)
		return err
	}
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	var userInfo updateUser.UserInfo
	userInfo.UserId = int32(myUser.ID)
	userInfo.Name = myUser.Name
	userInfo.Mobile = myUser.Mobile
	//userInfo.RealName = myUser.Real_name
	userInfo.IdCard = myUser.Id_card
	userInfo.AvatarUrl = myUser.Avatar_url
    rsp.Data = &userInfo

	return nil
}

func (e *GetUserInfo) MicroUpdateUser(ctx context.Context, req *updateUser.UpdateReq, rsp *updateUser.UpdateResp) error {

	return nil
}


