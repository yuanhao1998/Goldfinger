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
	"Goldfinger/utils/convert"
)

type UserServer struct {
	DataConn db.StringCache
}

func (u UserServer) Retrieve(ctx context.Context, req *userPB.RetrieveUserReq) (*userPB.RetrieveUserResp, error) {
	user := model.UMUser{}
	err := u.DataConn.RetrieveString(ctx, &user, userCacheKey, req.Id)
	if err != nil {
		return nil, err
	}

	result := userPB.RetrieveUserResp{}
	err = convert.StructToStructUseJson(user, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (u UserServer) Create(ctx context.Context, req *userPB.CreateUserReq) (*userPB.CreateUserResp, error) {
	user := model.UMUser{LoginName: req.LoginName, ShowName: req.ShowName, Password: req.Password, Mobile: req.Mobile, Desc: req.Desc, GroupId: req.GroupId}
	id, err := u.DataConn.CreateString(ctx, &user, userCacheKey)
	if err != nil {
		return nil, err
	}

	globals.Logger.Info("用户新增", id)
	globals.Logger.Debug("用户新增详情", req.String())

	return &userPB.CreateUserResp{Id: id}, nil
}

func (u UserServer) Delete(ctx context.Context, req *userPB.DeleteUserReq) (*userPB.DeleteUserResp, error) {
	id, err := u.DataConn.DeleteString(ctx, model.UMUser{Id: req.Id}, userCacheKey)
	if err != nil {
		return nil, err
	}

	globals.Logger.Info("用户删除", req.Id)
	return &userPB.DeleteUserResp{Id: id}, nil
}

func (u UserServer) Update(ctx context.Context, req *userPB.UpdateUserReq) (*userPB.UpdateUserResp, error) {
	user := model.UMUser{
		Id:        req.User.Id,
		LoginName: req.User.LoginName,
		ShowName:  req.User.ShowName,
		Password:  req.User.Password,
		Mobile:    req.User.Mobile,
		Desc:      req.User.Desc,
		GroupId:   req.User.GroupId,
	}

	id, err := u.DataConn.UpdateString(ctx, &user, userCacheKey)
	if err != nil {
		return nil, err
	}

	globals.Logger.Info("用户更新", req.User.Id)
	globals.Logger.Debug("用户更新详细信息：", req.User.String())

	return &userPB.UpdateUserResp{Id: id}, nil
}
