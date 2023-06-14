// @Create   : 2023/4/27 16:51
// @Author   : yaho
// @Remark   :

package model

type EntityCreateBaseModel struct { // 实体创建基础模型
	CreateUser  int64 `gorm:"not null;comment:创建此用户的用户id" redis:"CreateUser"`
	UpdateUser  int64 `gorm:"not null;comment:最后更新此用户的用户id" redis:"UpdateUser"`
	CreateGroup int64 `gorm:"not null;comment:创建此用户的用户组id" redis:"CreateGroup"`
	UpdateGroup int64 `gorm:"not null;comment:最后更新此用户的用户组id" redis:"UpdateGroup"`
	CreatedAt   int   `gorm:"comment:创建时间戳" redis:"CreateAt"`
	UpdatedAt   int   `gorm:"comment:更新时间戳" redis:"UpdateAt"`
}
