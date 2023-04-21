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

func CreateUserGroupHandler(query model.CreateUserGroupQueryModel, rc chan<- *userGroupPB.CreateResp, ok chan<- error) {
	userId, err := userGroupPB.NewUserGroupClient(globals.RPCClient).Create(context.Background(),
		&userGroupPB.CreateReq{Name: query.Name, ParentId: query.ParentId})
	if err != nil {
		ok <- err
		return
	}
	rc <- userId
}

func RetrieveUserGroupHandler(userGroupId int64, rc chan<- *userGroupPB.RetrieveResp, ok chan<- error) {
	retrieve, err := userGroupPB.NewUserGroupClient(globals.RPCClient).Retrieve(context.Background(), &userGroupPB.RetrieveReq{Id: userGroupId})
	if err != nil {
		ok <- err
		return
	}
	rc <- retrieve
}

func UpdateUserGroupHandler(query model.UpdateUserGroupQueryModel, rc chan<- *userGroupPB.UpdateResp, ok chan<- error) {
	userId, err := userGroupPB.NewUserGroupClient(globals.RPCClient).Update(
		context.Background(),
		&userGroupPB.UpdateReq{
			UserGroup: &userGroupPB.RetrieveResp{
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

func DeleteUserGroupHandler(userGroupId int64, rc chan<- *userGroupPB.DeleteResp, ok chan<- error) {
	userId, err := userGroupPB.NewUserGroupClient(globals.RPCClient).Delete(context.Background(), &userGroupPB.DeleteReq{Id: userGroupId})
	if err != nil {
		ok <- err
		return
	}
	rc <- userId
}
