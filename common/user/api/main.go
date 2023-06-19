// @Create   : 2023/3/17 23:07
// @Author   : yaho
// @Remark   :

package main

import (
	"fmt"

	"Goldfinger/common/user/api/src"
	_ "Goldfinger/common/user/config"
	"Goldfinger/common/user/globals"
	"Goldfinger/public/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()
	app.Use(middleware.CheckJWTAuth) // 注册鉴权中间件

	user.Route(app) // 注册url

	addr := fmt.Sprintf("%s:%s", userGlobals.RunConf.APIProject.Host, userGlobals.RunConf.APIProject.Port)
	panic(app.Run(addr))
}
