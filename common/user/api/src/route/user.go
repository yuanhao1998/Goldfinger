package route

import (
	"Goldfinger/common/user/api/src/view"
	"github.com/gin-gonic/gin"
)

func userV1(app *gin.RouterGroup) {
	v1 := app.Group("/v1")

	v1User := v1.Group("/user")
	{
		v1User.GET("/userGroup/:id", view.RetrieveGroupView)
		v1User.POST("/userGroup", view.CreateUserGroupView)
		v1User.PUT("/userGroup", view.UpdateGroupView)
		v1User.DELETE("/userGroup/:id", view.DeleteGroupView)

		v1User.POST("/user", view.CreateUserView)
		v1User.GET("/user/:id", view.RetrieveUserView)
	}
}
