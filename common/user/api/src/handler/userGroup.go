// @Create   : 2023/3/20 11:30
// @Author   : yaho
// @Remark   :

package handler

import (
	"context"

	"Goldfinger/common/user/api/src/model"
	"Goldfinger/common/user/globals"
	"Goldfinger/common/user/rpc/proto"
)

func CreateUserGroupHandler(ctx context.Context, query model.CreateUserGroupQueryModel, rc chan<- any, ok chan<- error) {
	groupId, err := userPB.NewUserGroupClient(globals.RPCClient).Create(ctx,
		&userPB.CreateUserGroupReq{Name: query.Name, ParentId: query.ParentId})
	if err != nil {
		ok <- err
		return
	}
	rc <- groupId
}

func RetrieveUserGroupHandler(ctx context.Context, userGroupId int64, rc chan<- any, ok chan<- error) {
	retrieve, err := userPB.NewUserGroupClient(globals.RPCClient).Retrieve(ctx, &userPB.RetrieveUserGroupReq{Id: userGroupId})
	if err != nil {
		ok <- err
		return
	}
	rc <- retrieve
}

func UpdateUserGroupHandler(ctx context.Context, query model.UpdateUserGroupQueryModel, rc chan<- any, ok chan<- error) {
	userId, err := userPB.NewUserGroupClient(globals.RPCClient).Update(
		ctx,
		&userPB.UpdateUserGroupReq{
			UserGroup: &userPB.RetrieveUserGroupResp{
				Id:       query.Id,
				Name:     query.Name,
				ParentId: query.ParentId,
			},
		},
	)
	if err != nil {
		ok <- err
		return
	}
	rc <- userId
}

func DeleteUserGroupHandler(ctx context.Context, userGroupId int64, rc chan<- any, ok chan<- error) {
	userId, err := userPB.NewUserGroupClient(globals.RPCClient).Delete(ctx, &userPB.DeleteUserGroupReq{Id: userGroupId})
	if err != nil {
		ok <- err
		return
	}
	rc <- userId
}
