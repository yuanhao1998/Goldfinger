// @Create   : 2023/4/23 09:49
// @Author   : yaho
// @Remark   :

package view

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"Goldfinger/common/user/api/src/handler"
	"Goldfinger/common/user/api/src/model"
	"Goldfinger/common/user/rpc/proto"
	"Goldfinger/config"
	"Goldfinger/errors"
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

	var resChan, errChane = make(chan *userPB.CreateUserResp), make(chan error)
	ctx, cancel := context.WithTimeout(c, config.APITimeOut)
	defer cancel()

	go handler.CreateUserHandler(query, resChan, errChane)

	for {
		select {
		case err := <-errChane:
			c.JSON(http.StatusInternalServerError, errors.NewParamsError(err.Error()).ErrorMap())
			return
		case res := <-resChan:
			c.JSON(http.StatusOK, gin.H{"userId": res.Id})
			return
		case <-ctx.Done(): //超时
			c.JSON(http.StatusInternalServerError, errors.NewTimeOutError("RPC请求超时").ErrorMap())
			return
		}
	}

}
