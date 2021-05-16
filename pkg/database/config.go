package database

import "time"

type Config struct {
	Name            string
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
