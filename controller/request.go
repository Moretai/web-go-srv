package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var CtxUserIdKey ="user_id"
var ErrorUserNotLogin = errors.New("用户未登录")

func getCurrentUser(c *gin.Context) (userID int64, err error) {
	 uid, ok := c.Get(CtxUserIdKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}

	if userID, ok = uid.(int64); !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}