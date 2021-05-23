package db

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var driverMap = map[string]func(c Conf) Driver{
	"mysql": NewMySqlDriver,
}

func New(c Conf) (*gorm.DB, error) {

	db, err := gorm.Open(selectDriver(c).Dialector())
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("open mysql failed. db name: %s, err: %+v", c.DbName, err))
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(c.MaxIdleConn)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(c.MaxOpenConn)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(c.ConnMaxLifeTime)
	return db, nil
}

func selectDriver(c Conf) Driver {
	newDriver, exist := driverMap[c.Driver]
	if !exist {
		newDriver = NewMySqlDriver
	}
	return newDriver(c)
}
