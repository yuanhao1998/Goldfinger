// @Create   : 2023/4/18 16:10
// @Author   : yaho
// @Remark   : 数值检查

package check

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"

	"Goldfinger/errors"
)

// Int64Check 检查数据是否符合int64
func Int64Check(data int64, c *gin.Context) bool {
	if data > math.MaxInt64 || data < math.MinInt64 {
		c.JSON(http.StatusBadRequest, errors.NewParamsError("Id超过最大/最小限制").ErrorMap())
		return false
	}
	return true
}
