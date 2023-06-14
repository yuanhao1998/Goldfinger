// @Create   : 2023/4/21 17:04
// @Author   : yaho
// @Remark   :

package handler

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"

	userGlobals "Goldfinger/common/user/globals"
	"Goldfinger/common/user/rpc/proto"
	"Goldfinger/common/user/rpc/src/model"
	"Goldfinger/errors"
	"Goldfinger/globals"
	"Goldfinger/public/db"
)

type LoginServer struct {
	DataConn db.StringCache
}

func GenerateJWTToken(user *model.UMUser) (string, error) { // 生成jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.Id,
		"showName":  user.ShowName,
		"loginName": user.LoginName,
		"groupId":   user.GroupId,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(userGlobals.RunConf.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (l LoginServer) Login(_ context.Context, req *userPB.LoginReq) (*userPB.LoginResp, error) {

	user := model.UMUser{}
	if l.DataConn.DbConn.Where("login_name = ? and is_del = 0").First(&user).RowsAffected == 0 { // 用户不存在
		return nil, errors.NewLoginError("登录失败")
	}

	if user.FailCount == 0 { // 账户锁定
		return nil, errors.NewLoginError("超过重试次数，账户已锁定")
	}

	if user.Password != req.Password { // 密码错误
		user.FailCount -= 1
		l.DataConn.DbConn.Save(&user)
		return nil, errors.NewLoginError("登陆失败")
	}

	token, err := GenerateJWTToken(&user)
	if err != nil {
		globals.Logger.Error("用户登陆时，生成jwt token失败：", err.Error())
		return nil, errors.NewLoginError("生成JWT Token失败")
	}

	user.LastLogin = time.Now().Unix()
	l.DataConn.DbConn.Save(&user)

	return &userPB.LoginResp{
		ShowName: user.ShowName,
		UserId:   user.Id,
		Token:    token,
	}, nil
}
