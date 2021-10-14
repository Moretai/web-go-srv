package controller

import (
	"net/http"
	"web_app/logic"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	// 1. 参数校验
	// 2.业务处理
	logic.SignUp()
	// 3.返回响应
	c.JSON(http.StatusOK, "ok")
}
