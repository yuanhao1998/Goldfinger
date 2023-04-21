// @Create   : 2023/3/21 16:28
// @Author   : yaho
// @Remark   :

package handler

import (
	proto "Goldfinger/public/proto"
	"context"
	"reflect"
	"strings"

	"Goldfinger/common/user/rpc/proto"
	"Goldfinger/common/user/rpc/src/model"
	"Goldfinger/public/db"
	"Goldfinger/utils/convert"
)

type UserGroupServer struct {
	DataConn db.StringCache[model.UMUserGroup]
}

func (u *UserGroupServer) Update(ctx context.Context, req *userGroupPB.UpdateReq) (*userGroupPB.UpdateResp, error) {
	userGroup := model.UMUserGroup{Id: req.UserGroup.Id, Name: req.UserGroup.Name, ParentId: req.UserGroup.ParentId}
	id, err := u.DataConn.UpdateString(ctx, &userGroup, "UMUserGroup")
	if err != nil {
		return nil, err
	}
	return &userGroupPB.UpdateResp{Id: id}, nil
}

func (u *UserGroupServer) Delete(ctx context.Context, req *userGroupPB.DeleteReq) (*userGroupPB.DeleteResp, error) {
	id, err := u.DataConn.DeleteString(ctx, "UMUserGroup", req.Id)
	if err != nil {
		return nil, err
	}
	return &userGroupPB.DeleteResp{Id: id}, nil
}

func (u *UserGroupServer) Query(ctx context.Context, req *userGroupPB.QueryReq) (*userGroupPB.QueryResp, error) {
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

	userGroups := userGroupPB.QueryResp{}
	u.DataConn.DbConn.Where(strings.Join(keySince, "and"), valueSince...).Find(&userGroups)
	return &userGroups, nil

}

func (u *UserGroupServer) Create(ctx context.Context, req *userGroupPB.CreateReq) (*userGroupPB.CreateResp, error) {
	userGroup := model.UMUserGroup{Name: req.Name, ParentId: req.ParentId}
	id, err := u.DataConn.CreateString(ctx, &userGroup, "UMUserGroup")
	if err != nil {
		return nil, err
	}
	return &userGroupPB.CreateResp{Id: id}, nil
}

func (u *UserGroupServer) Retrieve(ctx context.Context, req *userGroupPB.RetrieveReq) (*userGroupPB.RetrieveResp, error) {
	userGroup := model.UMUserGroup{}
	err := u.DataConn.RetrieveString(ctx, &userGroup, "UMUserGroup", req.Id)
	if err != nil {
		return nil, err
	}
	result := userGroupPB.RetrieveResp{}
	err = convert.JsonConvert(userGroup, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
