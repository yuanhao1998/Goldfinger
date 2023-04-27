// @Create   : 2023/4/23 09:22
// @Author   : yaho
// @Remark   :

package model

// LoginQueryModel 登陆接口参数模型
type LoginQueryModel struct {
	LoginName string `json:"loginName"` // 登陆名称
	Password  string `json:"password"`  // 密码
	CaptchaId string `json:"captchaId"` // 验证码Id
	Captcha   string `json:"captcha"`   // 验证码答案
}
