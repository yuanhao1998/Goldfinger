// @Create   : 2023/4/23 09:19
// @Author   : yaho
// @Remark   :

package view

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"Goldfinger/common/user/api/src/handler"
	"Goldfinger/common/user/api/src/model"
	"Goldfinger/common/user/rpc/proto"
	"Goldfinger/config"
	"Goldfinger/errors"
)

// CaptchaView 验证码生成接口
func CaptchaView(c *gin.Context) {
	id, b64s, err := handler.CaptchaHandler()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.NewCaptchaError(err.Error()).ErrorMap())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "b64s": b64s})
}

// LoginView 登陆接口
func LoginView(c *gin.Context) {

	var query model.LoginQueryModel
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewParamsError(err.Error()).ErrorMap())
		return
	}

	var resChan, errChane = make(chan *userPB.LoginResp), make(chan error)
	ctx, cancel := context.WithTimeout(c, config.APITimeOut)
	defer cancel()

	go handler.LoginHandler(query, resChan, errChane)

	for {
		select {
		case err := <-errChane:
			c.JSON(http.StatusInternalServerError, errors.NewParamsError(err.Error()).ErrorMap())
			return
		case res := <-resChan:
			c.Writer.Header().Set("token", res.Token)
			c.Writer.Header().Set("userId", strconv.FormatInt(res.UserId, 10))
			c.JSON(http.StatusOK, gin.H{"userName": res.ShowName})
			return
		case <-ctx.Done(): //超时
			c.JSON(http.StatusInternalServerError, errors.NewTimeOutError("RPC请求超时").ErrorMap())
			return
		}
	}

}
