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

func CreateUserGroupHandler(query model.CreateUserGroupQueryModel, rc chan<- *userPB.CreateUserGroupResp, ok chan<- error) {
	groupId, err := userPB.NewUserGroupClient(globals.RPCClient).Create(context.Background(),
		&userPB.CreateUserGroupReq{Name: query.Name, ParentId: query.ParentId})
	if err != nil {
		ok <- err
		return
	}
	rc <- groupId
}

func RetrieveUserGroupHandler(userGroupId int64, rc chan<- *userPB.RetrieveUserGroupResp, ok chan<- error) {
	retrieve, err := userPB.NewUserGroupClient(globals.RPCClient).Retrieve(context.Background(), &userPB.RetrieveUserGroupReq{Id: userGroupId})
	if err != nil {
		ok <- err
		return
	}
	rc <- retrieve
}

func UpdateUserGroupHandler(query model.UpdateUserGroupQueryModel, rc chan<- *userPB.UpdateUserGroupResp, ok chan<- error) {
	userId, err := userPB.NewUserGroupClient(globals.RPCClient).Update(
		context.Background(),
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

func DeleteUserGroupHandler(userGroupId int64, rc chan<- *userPB.DeleteUserGroupResp, ok chan<- error) {
	userId, err := userPB.NewUserGroupClient(globals.RPCClient).Delete(context.Background(), &userPB.DeleteUserGroupReq{Id: userGroupId})
	if err != nil {
		ok <- err
		return
	}
	rc <- userId
}
