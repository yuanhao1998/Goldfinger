// @Create   : 2023/3/20 20:20
// @Author   : yaho
// @Remark   : 服务注册

package src

import (
	"google.golang.org/grpc"

	"Goldfinger/common/user/globals"
	"Goldfinger/common/user/rpc/proto"
	"Goldfinger/common/user/rpc/src/handler"
	"Goldfinger/public/db"
)

func Register(server *grpc.Server) {
	userPB.RegisterUserGroupServer(server, &handler.UserGroupServer{DataConn: db.StringCache{DbConn: userGlobals.DBConn, CacheConn: userGlobals.CacheConn}})
	userPB.RegisterUserServer(server, &handler.UserServer{DataConn: db.StringCache{DbConn: userGlobals.DBConn, CacheConn: userGlobals.CacheConn}})
	userPB.RegisterLoginServer(server, &handler.LoginServer{DataConn: db.StringCache{DbConn: userGlobals.DBConn, CacheConn: userGlobals.CacheConn}})
}
