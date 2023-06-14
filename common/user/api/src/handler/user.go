// @Create   : 2023/4/23 10:00
// @Author   : yaho
// @Remark   :

package handler

import (
	"Goldfinger/common/user/api/src/model"
	"Goldfinger/common/user/globals"
	userPB "Goldfinger/common/user/rpc/proto"
	"context"
)

func CreateUserHandler(ctx context.Context, query model.CreateUserQueryModel, rc chan<- any, ok chan<- error) {
	userId, err := userPB.NewUserClient(globals.RPCClient).Create(ctx,
		&userPB.CreateUserReq{LoginName: query.LoginName, ShowName: query.ShowName, Password: query.Password, Mobile: query.Mobile, Desc: query.Desc, GroupId: query.GroupId})
	if err != nil {
		ok <- err
		return
	}
	rc <- userId
}
