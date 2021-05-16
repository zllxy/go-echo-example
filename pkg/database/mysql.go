package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Mysql *sql.DB

func NewMysql(conf *Config) {
	conn, err := gorm.Open(mysql.Open(getMysqlDsn(conf)), &gorm.Config{})
	if err != nil {
		log.Panicf("database connection failed. database name: %s, err: %+v", conf.Name, err)
	}
	Mysql, err = conn.DB()
	if err != nil {
		log.Panicf("database connection failed. driver name: mysql, err: %+v", err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	Mysql.SetMaxIdleConns(conf.MaxIdleConn)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	Mysql.SetMaxOpenConns(conf.MaxOpenConn)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	Mysql.SetConnMaxLifetime(conf.ConnMaxLifeTime)
}

func Close()  {
	Mysql.Close()
}

func getMysqlDsn(conf *Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=%s",
		conf.UserName,
		conf.Password,
		conf.Addr,
		conf.Name,
		conf.Charset,
		conf.ParseTime,
		conf.TimeZone,
	)
}
