// @Create   : 2023/5/25 15:18
// @Author   : yaho
// @Remark   :

package response

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"Goldfinger/config"
	"Goldfinger/errors"
	"Goldfinger/utils/convert"
)

func DefaultResponse(c *gin.Context, resChan chan any, errChane chan error) {
	ctx, cancel := context.WithTimeout(c, config.APITimeOut)
	defer cancel()

	for {
		select {
		case err := <-errChane:
			c.JSON(http.StatusInternalServerError, errors.NewParamsError(err.Error()).ErrorMap())
			return
		case res := <-resChan:
			c.JSON(http.StatusOK, gin.H(convert.StructToMapUseRef(res)))
			return
		case <-ctx.Done(): //超时
			c.JSON(http.StatusInternalServerError, errors.NewTimeOutError("RPC请求超时").ErrorMap())
			return
		}
	}
}
