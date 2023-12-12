// @Create   : 2023/3/21 16:28
// @Author   : yaho
// @Remark   :

package handler

import (
	"Goldfinger/errors"
	"context"
	"reflect"
	"strings"

	"Goldfinger/common/user/rpc/proto"
	"Goldfinger/common/user/rpc/src/model"
	"Goldfinger/globals"
	"Goldfinger/public/db"
	proto "Goldfinger/public/proto"
	"Goldfinger/utils/convert"
)

type UserGroupServer struct {
	DataConn db.StringCache
}

func (u *UserGroupServer) Create(ctx context.Context, req *userPB.CreateUserGroupReq) (*userPB.CreateUserGroupResp, error) {

	if u.DataConn.DbConn.Where("parent_id = ? AND is_del = 0 AND name = ?", req.ParentId, req.Name).RowsAffected != 0 {
		return nil, errors.NewNameRepeatError("用户组名称重复")
	}

	userGroup := model.UMUserGroup{Name: req.Name, ParentId: req.ParentId}
	id, err := u.DataConn.CreateString(ctx, &userGroup, userGroupCacheKey)
	if err != nil {
		return nil, err
	}

	globals.Logger.Info("用户组新增", id)
	globals.Logger.Debug("用户组新增详情", req.String())

	return &userPB.CreateUserGroupResp{Id: id}, nil
}

func (u *UserGroupServer) Update(ctx context.Context, req *userPB.UpdateUserGroupReq) (*userPB.UpdateUserGroupResp, error) {

	userGroup := model.UMUserGroup{Id: req.UserGroup.Id, Name: req.UserGroup.Name, ParentId: req.UserGroup.ParentId}
	id, err := u.DataConn.UpdateString(ctx, &userGroup, userGroupCacheKey)
	if err != nil {
		return nil, err
	}

	globals.Logger.Info("用户组更新", req.UserGroup.Id)
	globals.Logger.Debug("用户组更新详细信息：", req.UserGroup.String())

	return &userPB.UpdateUserGroupResp{Id: id}, nil
}

func (u *UserGroupServer) Delete(ctx context.Context, req *userPB.DeleteUserGroupReq) (*userPB.DeleteUserGroupResp, error) {

	id, err := u.DataConn.DeleteString(ctx, model.UMUserGroup{Id: req.Id}, userGroupCacheKey)
	if err != nil {
		return nil, err
	}

	globals.Logger.Info("用户组删除", req.Id)

	return &userPB.DeleteUserGroupResp{Id: id}, nil
}

func (u *UserGroupServer) Query(_ context.Context, req *userPB.QueryUserGroupReq) (*userPB.QueryUserGroupResp, error) {

	query := map[string]any{
		"id IN ?":                   req.Ids,
		"name Like ?":               "%" + strings.ToLower(req.Name) + "%",
		"parent_id IN ?":            req.ParentId,
		"is_admin = ?":              req.IsAdmin,
		"create_user = ?":           req.CreateUser,
		"update_user = ?":           req.UpdateUser,
		"create_group = ?":          req.CreateGroup,
		"update_group = ?":          req.UpdateGroup,
		"create_at BETWEEN ? AND ?": req.CreateAt,
		"update_at BETWEEN ? AND ?": req.UpdateAt,
	}

	var keySince []string // 查询条件切片
	var valueSince []any  // 查询值切片
	for k, v := range query {
		reflectData := reflect.ValueOf(v)
		if (reflectData.Kind() == reflect.Slice && reflectData.Len() != 0) ||
			(reflectData.Kind() == reflect.Int64 && v != 0) ||
			(reflectData.Kind() == reflect.Bool && v != proto.BoolSelectEnum_ALL) {

			keySince = append(keySince, k)
			valueSince = append(valueSince, v)
		}
	}

	userGroups := userPB.QueryUserGroupResp{}
	u.DataConn.DbConn.Where(strings.Join(keySince, "and"), valueSince...).Find(&userGroups)
	return &userGroups, nil

}

func (u *UserGroupServer) Retrieve(ctx context.Context, req *userPB.RetrieveUserGroupReq) (*userPB.RetrieveUserGroupResp, error) {

	userGroup := model.UMUserGroup{}
	err := u.DataConn.RetrieveString(ctx, &userGroup, userGroupCacheKey, req.Id)
	if err != nil {
		return nil, err
	}

	result := userPB.RetrieveUserGroupResp{}
	err = convert.StructToStructUseJson(userGroup, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
