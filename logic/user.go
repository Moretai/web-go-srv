package logic

import (
	"errors"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户存在与否
	if exist, err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	} else {
		if exist {
			return errors.New("user exist")
		}
	}
	// 生成UID
	userID := snowflake.GenID()

	u := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// 保存到数据库
	return mysql.InsertUser(u)
}
