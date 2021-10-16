package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"web_app/models"
)

const secret = "more"

func CheckUserExist(name string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int64
	if err := db.Get(&count, sqlStr, name); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func InsertUser(user *models.User) (err error) {
	user.Password = encryptPassword(user.Password)
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(oPwd string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPwd)))
}

func Login(u *models.User) (err error) {
	oPassword := u.Password
	sqlStr := `select user_id, username, password from user where username=?`
	err = db.Get(u, sqlStr, u.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	password := encryptPassword(oPassword)
	if password != u.Password {
		return ErrorInvalidPassword
	}
	return
}

func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, username from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	return
}
