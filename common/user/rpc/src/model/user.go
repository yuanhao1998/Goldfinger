// @Create   : 2023/4/23 10:15
// @Author   : yaho
// @Remark   :

package model

type UMUser struct {
	Id          int64  `gorm:"autoIncrement;primaryKey" redis:"Id"`
	LoginName   string `gorm:"type:varchar(50);not null;" redis:"LoginName"`
	ShowName    string `gorm:"type:varchar(50)not null" redis:"ShowName"`
	Password    string `gorm:"not null;"`
	Mobile      string `gorm:"type:varchar(50);not null;" redis:"Mobile"`
	Desc        string `redis:"Desc"`
	GroupId     int64  `gorm:"not null;" redis:"GroupId"`
	FailCount   int8   `gorm:"default:5"`
	IsDel       bool   `gorm:"type:bool;default:0;not null;" redis:"IsDel"`
	IsAdmin     bool   `gorm:"type:bool;default:0;not null;" redis:"IsAdmin"`
	LastLogin   int64
	CreateUser  int64 `gorm:"not null;" redis:"CreateUser"`
	UpdateUser  int64 `gorm:"not null;" redis:"UpdateUser"`
	CreateGroup int64 `gorm:"not null;" redis:"CreateGroup"`
	UpdateGroup int64 `gorm:"not null;" redis:"UpdateGroup"`
	CreatedAt   int   `redis:"CreateAt"`
	UpdatedAt   int   `redis:"UpdateAt"`
}
