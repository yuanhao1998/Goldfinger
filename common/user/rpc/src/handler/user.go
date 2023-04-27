// @Create   : 2023/4/23 10:33
// @Author   : yaho
// @Remark   :

package handler

import (
	"context"

	"Goldfinger/common/user/rpc/proto"
	"Goldfinger/common/user/rpc/src/model"
	"Goldfinger/globals"
	"Goldfinger/public/db"
)

type UserServer struct {
	DataConn db.StringCache
}

func (u UserServer) Create(ctx context.Context, req *userPB.CreateUserReq) (*userPB.CreateUserResp, error) {
	userGroup := model.UMUser{LoginName: req.LoginName, ShowName: req.ShowName, Password: req.Password, Mobile: req.Mobile, Desc: req.Desc, GroupId: req.GroupId}
	id, err := u.DataConn.CreateString(ctx, &userGroup, userCacheKey)
	if err != nil {
		return nil, err
	}

	globals.Logger.Info("用户新增", id)
	globals.Logger.Debug("用户新增详情", req.String())

	return &userPB.CreateUserResp{Id: id}, nil
}
