package database

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDialector(sqlDB *sql.DB) gorm.Dialector {
	return mysql.New(mysql.Config{Conn: sqlDB})
}
