// @Create   : 2023/6/19 10:52
// @Author   : yaho
// @Remark   : 权限相关中间件

package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"Goldfinger/common/user/globals"
	"Goldfinger/errors"
	"Goldfinger/globals"
)

type ParseUser struct { // 用于解析token的用户模型、需要增加jwt.StandardClaims字段
	id        int64
	showName  string
	loginName string
	groupId   int64
	exp       int64
	jwt.StandardClaims
}

// CheckJWTAuth jwt鉴权中间件，判断请求token是否过期、刷新临近过期token的时间
func CheckJWTAuth(c *gin.Context) {

	headerToken := c.Request.Header.Get("Token")
	if headerToken == "" {
		c.JSON(http.StatusUnauthorized, errors.NewAuthError("没有Token").ErrorMap())
		c.Abort()
		return
	}

	// 解析 JWT
	user := &ParseUser{}
	token, err := jwt.ParseWithClaims(headerToken, user, func(token *jwt.Token) (any, error) {
		return []byte(userGlobals.RunConf.SecretKey), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, errors.NewAuthError("Token解析错误").ErrorMap())
		c.Abort()
		return
	}

	// 验证 JWT 是否有效
	if !token.Valid {
		c.JSON(http.StatusUnauthorized, errors.NewAuthError("Token失效，请重新登录").ErrorMap())
		c.Abort()
		return
	} else if user.exp > time.Now().Add(time.Hour*6).Unix() {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":        user.Id,
			"showName":  user.showName,
			"loginName": user.loginName,
			"groupId":   user.groupId,
			"exp":       time.Now().Add(time.Hour * 12).Unix(),
		})

		// 新的token生成失败，但仍然可以使用原来的有效token，故此处只进行日志记录
		if newToken, err := token.SignedString([]byte(userGlobals.RunConf.SecretKey)); err != nil {
			globals.Logger.Error("token鉴权时发现有效期小于6小时，尝试刷新token时生成新的token失败：" + err.Error())
		} else {
			c.Header("NewToken", newToken)
		}
	}

	c.Next()
}
