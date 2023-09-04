// @Create   : 2023/4/23 09:19
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

	var resChan, errChan = make(chan any), make(chan error)
	go handler.LoginHandler(c, query, resChan, errChan)
	response.HeadersResponse(c, resChan, errChan)

}
