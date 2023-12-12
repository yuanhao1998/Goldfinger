// @Create   : 2023/3/17 23:07
// @Author   : yaho
// @Remark   :

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"Goldfinger/common/user/api/src/route"
	_ "Goldfinger/common/user/config" // 初始化配置，不要删除
	"Goldfinger/common/user/globals"
	"Goldfinger/public/middleware"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	baseApp := gin.Default()
	app := baseApp.Group("/api/user")
	noAuthApp := baseApp.Group("/api/user")

	app.Use(middleware.CheckJWTAuth) // 注册鉴权中间件

	route.Route(app, noAuthApp) // 注册url

	addr := fmt.Sprintf("%s:%s", userGlobals.RunConf.APIProject.Host, userGlobals.RunConf.APIProject.Port)
	panic(baseApp.Run(addr))
}
