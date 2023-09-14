// @Create   : 2023/3/16 20:54
// @Author   : yaho
// @Remark   : 用户模块路由

package route

import (
	"github.com/gin-gonic/gin"
)

func Route(app, noAuthApp *gin.RouterGroup) {
	userRoute(app)        // 用户
	loginRoute(noAuthApp) // 登录
}

func userRoute(app *gin.RouterGroup) {
	userV1(app)
}

func loginRoute(app *gin.RouterGroup) {
	loginV1(app)
}
