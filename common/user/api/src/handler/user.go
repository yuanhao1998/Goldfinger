// @Create   : 2023/4/23 10:00
// @Author   : yaho
// @Remark   :

package handler

import (
	"Goldfinger/common/user/api/src/model"
	"Goldfinger/common/user/globals"
	"Goldfinger/common/user/rpc/proto"
	"context"
)

func CreateUserHandler(ctx context.Context, query model.CreateUserQueryModel, rc chan<- any, ok chan<- error) {

	userId, err := userPB.NewUserClient(userGlobals.RPCClient).Create(ctx,
		&userPB.CreateUserReq{LoginName: query.LoginName, ShowName: query.ShowName, Password: query.Password, Mobile: query.Mobile, Desc: query.Desc, GroupId: query.GroupId})

	if err != nil {
		ok <- err
		return
	}

	rc <- userId
}

func RetrieveUserHandler(ctx context.Context, id int64, rc chan<- any, ok chan<- error) {

	retrieve, err := userPB.NewUserClient(userGlobals.RPCClient).Retrieve(ctx,
		&userPB.RetrieveUserReq{Id: id})

	if err != nil {
		ok <- err
		return
	}

	rc <- retrieve
}
