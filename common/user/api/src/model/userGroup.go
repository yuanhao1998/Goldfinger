// @Create   : 2023/3/20 10:02
// @Author   : yaho
// @Remark   :

package model

// CreateUserGroupQueryModel 创建用户组参数模型
type CreateUserGroupQueryModel struct {
	Name     string `json:"name"`
	ParentId int64  `json:"parentId"`
}

// UpdateUserGroupQueryModel 更新用户组参数模型
type UpdateUserGroupQueryModel struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	ParentId int64  `json:"parentId"`
}
