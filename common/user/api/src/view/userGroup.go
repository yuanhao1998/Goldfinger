// @Create   : 2023/3/17 17:42
// @Author   : yaho
// @Remark   :

package view

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

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

	var resChan, errChane = make(chan *userGroupPB.CreateResp), make(chan error)
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
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

	var resChan, errChane = make(chan *userGroupPB.RetrieveResp), make(chan error)
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
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

	var resChan, errChane = make(chan *userGroupPB.UpdateResp), make(chan error)
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
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

	var resChan, errChane = make(chan *userGroupPB.DeleteResp), make(chan error)
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
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
