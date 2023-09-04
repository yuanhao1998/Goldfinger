// @Create   : 2023/4/23 09:49
// @Author   : yaho
// @Remark   :

package view

import (
	"net/http"
	"strconv"

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

	var resChan, errChan = make(chan any), make(chan error)
	go handler.CreateUserHandler(c, query, resChan, errChan)
	response.DefaultResponse(c, resChan, errChan)
}

func RetrieveUserView(c *gin.Context) {

	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewParamsError("id无法转为int类型").ErrorMap())
	}

	var resChan, errChan = make(chan any), make(chan error)
	go handler.RetrieveUserHandler(c, intId, resChan, errChan)
	response.DefaultResponse(c, resChan, errChan)
}
