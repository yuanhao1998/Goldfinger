// @Create   : 2023/3/21 16:17
// @Author   : yaho
// @Remark   :

package model

type UMUserGroup struct {
	Id       int64  `gorm:"autoIncrement;primaryKey;comment:用户id" redis:"Id"`
	Name     string `gorm:"type:varchar(50);not null;comment:用户名称" redis:"Name"`
	ParentId int64  `gorm:"not null;comment:用户所属组id" redis:"ParentId"`
	IsDel    bool   `gorm:"type:bool;default:0;not null;comment:是否删除" redis:"IsDel"`
	IsAdmin  bool   `gorm:"type:bool;default:0;not null;comment:是否管理员" redis:"IsAdmin"`
	Desc     string `gorm:"comment:备注" redis:"Desc"`
	EntityCreateBaseModel
}
