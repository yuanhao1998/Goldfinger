// @Create   : 2023/4/17 17:56
// @Author   : yaho
// @Remark   :

package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	_ "Goldfinger/common/user/config" // 初始化配置，不要删除
	"Goldfinger/common/user/globals"
	"Goldfinger/common/user/rpc/src"
	"Goldfinger/config"
	"Goldfinger/globals"
)

func main() {

	// 日志初始化
	globals.Logger = config.InitLog(userGlobals.RunConf.RPCLog.Level, userGlobals.RunConf.RPCLog.Path)

	// grpc注册
	grpcServer := grpc.NewServer()
	src.Register(grpcServer)

	addr := fmt.Sprintf("%s:%s", userGlobals.RunConf.RPCProject.Host, userGlobals.RunConf.RPCProject.Port)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic("监听端口失败:" + err.Error())
	}

	globals.Logger.Info("启动User服务成功，监听RPC端口：", addr)

	err = grpcServer.Serve(listen)
	if err != nil {
		panic("启动服务失败：" + err.Error())
	}
}
