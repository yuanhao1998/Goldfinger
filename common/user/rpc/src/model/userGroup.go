// @Create   : 2023/3/21 16:17
// @Author   : yaho
// @Remark   :

package model

type UMUserGroup struct {
	Id          int64  `gorm:"autoIncrement;primaryKey" redis:"Id"`
	Name        string `gorm:"type:varchar(50);not null;" redis:"Name"`
	ParentId    int64  `gorm:"not null" redis:"ParentId"`
	IsDel       bool   `gorm:"type:bool;default:0;not null;" redis:"IsDel"`
	IsAdmin     bool   `gorm:"type:bool;default:0;not null;" redis:"IsAdmin"`
	Desc        string `gorm:"type:char;" redis:"Desc"`
	CreateUser  int64  `gorm:"not null;" redis:"CreateUser"`
	UpdateUser  int64  `gorm:"not null;" redis:"UpdateUser"`
	CreateGroup int64  `gorm:"not null;" redis:"CreateGroup"`
	UpdateGroup int64  `gorm:"not null;" redis:"UpdateGroup"`
	CreatedAt   int    `redis:"CreateAt"`
	UpdatedAt   int    `redis:"UpdateAt"`
}
