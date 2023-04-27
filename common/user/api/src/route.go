// @Create   : 2023/3/16 20:54
// @Author   : yaho
// @Remark   : 用户模块路由

package user

import (
	"github.com/gin-gonic/gin"

	"Goldfinger/common/user/api/src/view"
)

func Route(app *gin.Engine) {

	v1 := app.Group("/api/v1")

	v1User := v1.Group("/user")
	{
		v1User.GET("/userGroup/:id", view.RetrieveGroupView)
		v1User.POST("/userGroup", view.CreateUserGroupView)
		v1User.PUT("/userGroup", view.UpdateGroupView)
		v1User.DELETE("/userGroup/:id", view.DeleteGroupView)

		v1User.POST("/user", view.CreateUserView)
	}

	v1Login := v1.Group("/login")
	{
		v1Login.GET("/captcha", view.CaptchaView)
		v1Login.POST("/login", view.LoginView)
	}

}
