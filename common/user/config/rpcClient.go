// @Create   : 2023/3/23 15:10
// @Author   : yaho
// @Remark   :

package config

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"Goldfinger/common/user/globals"
)

func init() {
	// 连接到 gRPC 服务器
	target := fmt.Sprintf("%s:%s", globals.RunConf.RPCProject.Host, globals.RunConf.RPCProject.Port)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("连接RPC服务失败：" + err.Error())
	}

	globals.RPCClient = conn

}
