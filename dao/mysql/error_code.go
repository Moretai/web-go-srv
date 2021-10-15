package mysql

import "errors"
var  (
	ErrorInvalidID = errors.New("无效的ID")
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
)