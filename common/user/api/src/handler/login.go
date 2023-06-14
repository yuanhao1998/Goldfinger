// @Create   : 2023/4/25 14:58
// @Author   : yaho
// @Remark   :

package handler

import (
	"Goldfinger/common/user/api/src/model"
	globals2 "Goldfinger/common/user/globals"
	userPB "Goldfinger/common/user/rpc/proto"
	"Goldfinger/globals"
	"context"
	"github.com/mojocn/base64Captcha"
	"image/color"

	"Goldfinger/errors"
)

var captchaStore = base64Captcha.DefaultMemStore // 验证码存储空间

func CaptchaHandler() (string, string, error) {

	var fontsStorage base64Captcha.FontsStorage
	var fonts []string
	driver := base64Captcha.NewDriverMath(
		70, 240, 0,
		base64Captcha.OptionShowHollowLine,
		&color.RGBA{R: 144, G: 238, B: 144, A: 10},
		fontsStorage, fonts,
	)

	captcha := base64Captcha.NewCaptcha(driver, captchaStore)
	id, b64s, err := captcha.Generate()

	if err != nil {
		globals.Logger.Error("验证码生成出错：" + err.Error())
		return "", "", errors.NewCaptchaError("验证码生成出错")
	}

	return id, b64s, nil
}

func LoginHandler(ctx context.Context, query model.LoginQueryModel, rc chan<- any, ok chan<- error) {

	if len(query.CaptchaId) == 0 || len(query.Captcha) == 0 || !captchaStore.Verify(query.CaptchaId, query.Captcha, true) {
		ok <- errors.NewCaptchaError("验证码计算错误")
		return
	}

	token, err := userPB.NewLoginClient(globals2.RPCClient).Login(ctx, &userPB.LoginReq{LoginName: query.LoginName, Password: query.Password})
	if err != nil {
		ok <- err
		return
	}

	rc <- token
}
