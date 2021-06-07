package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlDriver struct {
	username  string
	password  string
	addr      string
	dbName    string
	charset   string
	parseTime bool
	timeZone  string
}

func NewMySqlDriver(c Conf) Driver {
	return &MySqlDriver{
		username:  c.UserName,
		password:  c.Password,
		addr:      c.Addr,
		dbName:    c.DbName,
		charset:   c.Charset,
		parseTime: c.ParseTime,
		timeZone:  c.TimeZone,
	}
}

func (d *MySqlDriver) ParseDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=%s",
		d.username,
		d.password,
		d.addr,
		d.dbName,
		d.charset,
		d.parseTime,
		d.timeZone,
	)
}

func (d *MySqlDriver) Dialector() gorm.Dialector {
	return mysql.Open(d.ParseDSN())
}
