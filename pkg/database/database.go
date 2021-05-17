package database

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"log"
)

var dialector = map[string]func(sqlDB *sql.DB) gorm.Dialector{
	"mysql": NewMysqlDialector,
}

type Database struct {
	conf *Config
	Dsn  string
	Conn *sql.DB
	DB   *gorm.DB
}

func NewDatabase(conf *Config) *Database {
	db := &Database{conf: conf}
	db.setDsn()
	db.setConnPool()
	db.setDB()
	return db
}

func (db *Database) setDsn() {
	db.Dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=%s",
		db.conf.UserName,
		db.conf.Password,
		db.conf.Addr,
		db.conf.Name,
		db.conf.Charset,
		db.conf.ParseTime,
		db.conf.TimeZone,
	)
}

func (db *Database) setConnPool() {
	sqlDB, err := sql.Open(db.conf.Driver, db.Dsn)
	if err != nil {
		log.Panicf("open mysql failed. database name: %s, err: %+v", db.conf.Name, err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(db.conf.MaxIdleConn)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(db.conf.MaxOpenConn)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(db.conf.ConnMaxLifeTime)

	db.Conn = sqlDB
}

func (db *Database) setDB() {
	var dial func(sqlDB *sql.DB) gorm.Dialector
	dial, exist := dialector[db.conf.Driver]
	if !exist {
		dial = NewMysqlDialector
	}
	DB, err := gorm.Open(dial(db.Conn))
	if err != nil {
		log.Panicf("database connection failed. database name: %s, err: %+v", db.conf.Name, err)
	}
	db.DB = DB
}

func (db Database) Close() {
	db.Conn.Close()
}
