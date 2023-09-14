package route

import (
	"Goldfinger/common/user/api/src/view"
	"github.com/gin-gonic/gin"
)

func loginV1(app *gin.RouterGroup) {
	v1 := app.Group("/v1")

	v1Login := v1.Group("/login")
	v1Login.Use() // 登录相关，不使用全局中间件鉴权
	{
		v1Login.GET("/captcha", view.CaptchaView)
		v1Login.POST("/login", view.LoginView)
	}
}
