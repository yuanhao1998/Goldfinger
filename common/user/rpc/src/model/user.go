// @Create   : 2023/4/23 10:15
// @Author   : yaho
// @Remark   :

package model

type UMUser struct {
	Id        int64  `gorm:"autoIncrement;primaryKey;comment:用户id" redis:"Id"`
	LoginName string `gorm:"type:varchar(50);not null;comment:用户登录名称" redis:"LoginName"`
	ShowName  string `gorm:"type:varchar(50)not null;comment:用户展示名称" redis:"ShowName"`
	Password  string `gorm:"not null;comment:密码"`
	Mobile    string `gorm:"type:varchar(50);not null;comment:手机号" redis:"Mobile"`
	Desc      string `gorm:"comment:备注" redis:"Desc"`
	GroupId   int64  `gorm:"not null;comment:用户组id" redis:"GroupId"`
	FailCount int8   `gorm:"default:5;comment:失败重试次数"`
	IsDel     bool   `gorm:"type:bool;default:0;not null;comment:是否删除" redis:"IsDel"`
	IsAdmin   bool   `gorm:"type:bool;default:0;not null;comment:是否管理" redis:"IsAdmin"`
	LastLogin int64  `gorm:"comment:最近登录时间"`
	EntityCreateBaseModel
}
