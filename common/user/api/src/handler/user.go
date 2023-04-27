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

func CreateUserHandler(query model.CreateUserQueryModel, rc chan<- *userPB.CreateUserResp, ok chan<- error) {
	userId, err := userPB.NewUserClient(globals.RPCClient).Create(context.Background(),
		&userPB.CreateUserReq{LoginName: query.LoginName, ShowName: query.ShowName, Password: query.Password, Mobile: query.Mobile, Desc: query.Desc, GroupId: query.GroupId})
	if err != nil {
		ok <- err
		return
	}
	rc <- userId
}
