// @Create   : 2023/3/20 20:20
// @Author   : yaho
// @Remark   : 服务注册

package src

import (
	"google.golang.org/grpc"

	"Goldfinger/common/user/globals"
	"Goldfinger/common/user/rpc/proto"
	"Goldfinger/common/user/rpc/src/handler"
	"Goldfinger/common/user/rpc/src/model"
	"Goldfinger/public/db"
)

func Register(server *grpc.Server) {
	userGroupPB.RegisterUserGroupServer(server, &handler.UserGroupServer{DataConn: db.StringCache[model.UMUserGroup]{DbConn: globals.DBConn, CacheConn: globals.CacheConn}})
}
