// @Create   : 2023/4/23 09:49
// @Author   : yaho
// @Remark   :

package view

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"Goldfinger/common/user/api/src/handler"
	"Goldfinger/common/user/api/src/model"
	"Goldfinger/errors"
	"Goldfinger/public/view/response"
)

func CreateUserView(c *gin.Context) {

	var query model.CreateUserQueryModel
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewParamsError(err.Error()).ErrorMap())
		return
	}

	if query.Password != query.ConfirmPassword {
		c.JSON(http.StatusOK, errors.NewParamsError("两次密码输入不一致").ErrorMap())
	}

	var resChan, errChane = make(chan any), make(chan error)
	go handler.CreateUserHandler(c, query, resChan, errChane)
	response.DefaultResponse(c, resChan, errChane)
}
