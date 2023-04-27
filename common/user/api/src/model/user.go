// @Create   : 2023/4/23 09:56
// @Author   : yaho
// @Remark   :

package model

type CreateUserQueryModel struct {
	LoginName       string `json:"loginName"`       // 登陆名称
	ShowName        string `json:"showName"`        // 用户名称
	Password        string `json:"password"`        // 密码
	ConfirmPassword string `json:"confirmPassword"` // 确认密码
	Mobile          string `json:"mobile"`          // 手机号
	Desc            string `json:"desc"`            // 备注
	GroupId         int64  `json:"groupId"`         // 所属用户组
}
