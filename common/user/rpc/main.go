// @Create   : 2023/4/17 17:56
// @Author   : yaho
// @Remark   :

package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	userConfig "Goldfinger/common/user/config"
	userGlobals "Goldfinger/common/user/globals"
	"Goldfinger/common/user/rpc/src"
	"Goldfinger/common/user/rpc/src/model"
	"Goldfinger/config"
	"Goldfinger/globals"
)

func main() {

	globals.Logger = config.InitLog(userGlobals.RunConf.RPCLog.Level, userGlobals.RunConf.RPCLog.Path)

	dbConn := userConfig.InitDB(userGlobals.RunConf)
	model.CreateTable(dbConn)   // 自动建表
	userGlobals.DBConn = dbConn // 全局数据库连接

	sessionConn := userConfig.InitCache(userGlobals.RunConf, userGlobals.RunConf.Redis.SessionDB)
	userGlobals.SessionConn = sessionConn // 用户状态
	cacheConn := userConfig.InitCache(userGlobals.RunConf, userGlobals.RunConf.Redis.CacheDB)
	userGlobals.CacheConn = cacheConn // 全局缓存连接

	grpcServer := grpc.NewServer()
	src.Register(grpcServer)

	addr := fmt.Sprintf("%s:%s", userGlobals.RunConf.RPCProject.Host, userGlobals.RunConf.RPCProject.Port)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic("监听端口失败:" + err.Error())
	}

	globals.Logger.Error("启动User服务成功，监听RPC端口：", addr)

	err = grpcServer.Serve(listen)
	if err != nil {
		panic("启动服务失败：" + err.Error())
	}
}
