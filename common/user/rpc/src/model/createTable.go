// @Create   : 2023/3/21 17:00
// @Author   : yaho
// @Remark   :

package model

import (
	"gorm.io/gorm"
)

func CreateTable(conn *gorm.DB) {
	err := conn.AutoMigrate(&UMUserGroup{})
	if err != nil {
		panic("自动建表失败：" + err.Error())
	}
}
