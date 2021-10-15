package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户存在与否
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
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

func Login(p *models.ParamLogin) (token string, err error) {
	u := &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	// u 是指针，能拿到 id
	if err := mysql.Login(u); err != nil {
		return "", err
	}
	// 生成token
	return jwt.GenToken(u.UserID, u.Username)
}
