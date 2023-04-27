// @Create   : 2023/3/17 17:42
// @Author   : yaho
// @Remark   :

package view

import (
	"Goldfinger/config"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"Goldfinger/common/user/api/src/handler"
	"Goldfinger/common/user/api/src/model"
	"Goldfinger/common/user/rpc/proto"
	"Goldfinger/errors"
	"Goldfinger/public/view/check"
)

func CreateUserGroupView(c *gin.Context) {

	var query model.CreateUserGroupQueryModel
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewParamsError(err.Error()).ErrorMap())
		return
	}

	var resChan, errChane = make(chan *userPB.CreateUserGroupResp), make(chan error)
	ctx, cancel := context.WithTimeout(c, config.APITimeOut)
	defer cancel()

	go handler.CreateUserGroupHandler(query, resChan, errChane)

	for {
		select {
		case err := <-errChane:
			c.JSON(http.StatusInternalServerError, errors.NewParamsError(err.Error()).ErrorMap())
			return
		case res := <-resChan:
			c.JSON(http.StatusOK, gin.H{"userGroupId": res.Id})
			return
		case <-ctx.Done(): //超时
			c.JSON(http.StatusInternalServerError, errors.NewTimeOutError("RPC请求超时").ErrorMap())
			return
		}
	}
}

func RetrieveGroupView(c *gin.Context) {

	userGroupId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewParamsError(err.Error()).ErrorMap())
		return
	}
	if !check.Int64Check(userGroupId, c) {
		return
	}

	var resChan, errChane = make(chan *userPB.RetrieveUserGroupResp), make(chan error)
	ctx, cancel := context.WithTimeout(c, config.APITimeOut)
	defer cancel()

	go handler.RetrieveUserGroupHandler(userGroupId, resChan, errChane)

	for {
		select {
		case err := <-errChane:
			c.JSON(http.StatusInternalServerError, errors.NewParamsError(err.Error()).ErrorMap())
			return
		case res := <-resChan:
			c.JSON(http.StatusOK, gin.H{"name": res.Name, "parentId": res.ParentId, "desc": res.Desc, "isAdmin": res.IsAdmin})
			return
		case <-ctx.Done(): //超时
			c.JSON(http.StatusInternalServerError, errors.NewTimeOutError("RPC请求超时").ErrorMap())
			return
		}
	}

}

func UpdateGroupView(c *gin.Context) {

	var query model.UpdateUserGroupQueryModel
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewParamsError(err.Error()).ErrorMap())
		return
	}
	if !check.Int64Check(query.Id, c) {
		return
	}

	var resChan, errChane = make(chan *userPB.UpdateUserGroupResp), make(chan error)
	ctx, cancel := context.WithTimeout(c, config.APITimeOut)
	defer cancel()

	go handler.UpdateUserGroupHandler(query, resChan, errChane)

	for {
		select {
		case err := <-errChane:
			c.JSON(http.StatusInternalServerError, errors.NewParamsError(err.Error()).ErrorMap())
			return
		case res := <-resChan:
			c.JSON(http.StatusOK, gin.H{"userGroupId": res.Id})
			return
		case <-ctx.Done(): //超时
			c.JSON(http.StatusInternalServerError, errors.NewTimeOutError("RPC请求超时").ErrorMap())
			return
		}
	}
}

func DeleteGroupView(c *gin.Context) {
	userGroupId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewParamsError(err.Error()).ErrorMap())
		return
	}
	if !check.Int64Check(userGroupId, c) {
		return
	}

	var resChan, errChane = make(chan *userPB.DeleteUserGroupResp), make(chan error)
	ctx, cancel := context.WithTimeout(c, config.APITimeOut)
	defer cancel()

	go handler.DeleteUserGroupHandler(userGroupId, resChan, errChane)

	for {
		select {
		case err := <-errChane:
			c.JSON(http.StatusInternalServerError, errors.NewParamsError(err.Error()).ErrorMap())
			return
		case res := <-resChan:
			c.JSON(http.StatusOK, gin.H{"userGroupId": res.Id})
			return
		case <-ctx.Done(): //超时
			c.JSON(http.StatusInternalServerError, errors.NewTimeOutError("RPC请求超时").ErrorMap())
			return
		}
	}
}
