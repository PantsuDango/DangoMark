package controller

import (
	"DangoMark/model/tables"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 新建职位
func (Controller Controller) Init(ctx *gin.Context, user tables.User) {

	JSONSuccess(ctx, http.StatusOK, user)
}
