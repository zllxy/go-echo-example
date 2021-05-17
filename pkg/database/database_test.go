package database

import (
	"testing"
	"time"
)

type User struct {
	Id       int    `gorm.id`
	Name     string `gorm.name`
	Age      int    `gorm.age`
	Birthday int64  `gorm.birthday`
}

func TestNewDatabase(t *testing.T) {
	db := NewDatabase(getConfig())
	defer db.Close()
	db.DB.Create(&User{
		Name:     "zzz",
		Age:      18,
		Birthday: time.Now().Unix(),
	})
}

func TestSelect(t *testing.T) {
	db := NewDatabase(getConfig())
	defer db.Close()
	var user User
	db.DB.First(&user)
	t.Log(user)
}

func getConfig() *Config {
	return &Config{
		Driver:          "mysql",
		Name:            "test",
		Addr:            "127.0.0.1:3306",
		UserName:        "root",
		Password:        "root",
		ShowLog:         true,
		MaxIdleConn:     10,
		MaxOpenConn:     60,
		ConnMaxLifeTime: 60,
		TimeZone:        "Asia%2fShanghai",
		Charset:         "utf8mb4",
		ParseTime:       true,
	}
}
