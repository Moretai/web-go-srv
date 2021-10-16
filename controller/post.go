package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/logic"
	"web_app/models"
)

func CreatePostHandler(c *gin.Context) {

	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("logic.CreatePost ShouldBindJSON failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
	}
	p.AuthorID = userID
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

func GetPostById(c *gin.Context) {
	idStr := c.Param("id")
	pid, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("GetPostById invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetPostById(pid)

	if err != nil {
		zap.L().Error("GetPostById failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}
