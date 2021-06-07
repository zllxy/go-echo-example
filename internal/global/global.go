package global

import (
	"go-echo-example/pkg/conf"
	"go-echo-example/pkg/log"
	"gorm.io/gorm"
)

var (
	Config *conf.Conf
	Logger *log.Logger
	DB     *gorm.DB
)
