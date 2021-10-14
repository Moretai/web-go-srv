package logic

import (
	"web_app/dao/mysql"
	"web_app/pkg/snowflake"
)

func SignUp() {
	// 判断用户存在与否
	mysql.QueryUserByUsername()
	// 生成UID
	snowflake.GenID()
	// 保存到数据库
	mysql.InsertUser()
}
