// @Create   : 2023/3/23 15:10
// @Author   : yaho
// @Remark   :

package userConfig

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"Goldfinger/common/user/globals"
)

func init() {
	// 连接到 gRPC 服务器
	target := fmt.Sprintf("%s:%s", userGlobals.RunConf.RPCProject.Host, userGlobals.RunConf.RPCProject.Port)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("连接RPC服务失败：" + err.Error())
	}

	userGlobals.RPCClient = conn

}
