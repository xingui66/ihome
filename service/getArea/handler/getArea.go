package handler

import (
	"context"


	getArea "ihomegit/ihome/service/getArea/proto/getArea"
	"ihomegit/ihome/service/getArea/model"
	"ihomegit/ihome/service/getArea/utils"
)

type GetArea struct{}

func (e *GetArea) MicroGetArea(ctx context.Context, req *getArea.Request, rsp *getArea.Response) error {
	//获取数据并返回给调用者
	areas ,err := model.QueryArea()
	if err != nil {
		rsp.Errno=utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	for _,v := range areas{
		var areaInfo getArea.AreaInfo
		areaInfo.Aid = int32(v.Id)
		areaInfo.Aname = v.Name

		rsp.Data = append(rsp.Data,&areaInfo)
	}

	return nil
}