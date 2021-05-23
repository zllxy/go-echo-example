package db

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

func TestInit(t *testing.T) {
	db, _ := New(getConfig())

	db.Create(&User{
		Name:     "zzz",
		Age:      18,
		Birthday: time.Now().Unix(),
	})
}

func TestSelect(t *testing.T) {
	db, _ := New(getConfig())
	var user User
	db.First(&user)
	t.Log(user)
}

func getConfig() *Conf {
	return &Conf{
		Driver:          "mysql",
		DbName:          "test",
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
