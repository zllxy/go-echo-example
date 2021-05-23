package db

import "time"

type Conf struct {
	Driver          string
	DbName          string
	Addr            string
	UserName        string
	Password        string
	ShowLog         bool
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifeTime time.Duration
	TimeZone        string
	Charset         string
	ParseTime       bool
}
