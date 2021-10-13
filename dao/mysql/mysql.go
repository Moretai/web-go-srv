package mysql

import (
	"fmt"
	"web_app/settings"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(config *settings.MySQLConfig) (err error) {
	// DSN:Data Source Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Dbname,
	)

	if db, err = sqlx.Connect("mysql", dsn); err != nil {
		zap.L().Error("init mysql failed", zap.Error(err))
		return
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	return
}

func Close() {
	_ = db.Close()
}
