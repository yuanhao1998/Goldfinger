// @Create   : 2023/3/17 17:42
// @Author   : yaho
// @Remark   :

package view

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"Goldfinger/common/user/api/src/handler"
	"Goldfinger/common/user/api/src/model"
	"Goldfinger/errors"
	"Goldfinger/public/view/check"
	"Goldfinger/public/view/response"
)

func CreateUserGroupView(c *gin.Context) {

	var query model.CreateUserGroupQueryModel
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewParamsError(err.Error()).ErrorMap())
		return
	}

	var resChan, errChan = make(chan any), make(chan error)
	go handler.CreateUserGroupHandler(c, query, resChan, errChan)
	response.DefaultResponse(c, resChan, errChan)

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

	var resChan, errChan = make(chan any), make(chan error)
	go handler.RetrieveUserGroupHandler(c, userGroupId, resChan, errChan)
	response.DefaultResponse(c, resChan, errChan)
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

	var resChan, errChan = make(chan any), make(chan error)
	go handler.UpdateUserGroupHandler(c, query, resChan, errChan)
	response.DefaultResponse(c, resChan, errChan)

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

	var resChan, errChan = make(chan any), make(chan error)
	go handler.DeleteUserGroupHandler(c, userGroupId, resChan, errChan)
	response.DefaultResponse(c, resChan, errChan)
}
