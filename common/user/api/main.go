// @Create   : 2023/3/17 23:07
// @Author   : yaho
// @Remark   :

package main

import (
	"fmt"

	"Goldfinger/common/user/api/src"
	_ "Goldfinger/common/user/config"
	"Goldfinger/common/user/globals"
	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()

	user.Route(app) // 注册url

	addr := fmt.Sprintf("%s:%s", globals.RunConf.APIProject.Host, globals.RunConf.APIProject.Port)
	panic(app.Run(addr))
}
