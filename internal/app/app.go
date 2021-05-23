package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	"go-echo-example/internal/global"
	"go-echo-example/pkg/conf"
	"go-echo-example/pkg/db"
	Logger "go-echo-example/pkg/log"
	"log"
)

const Path = "./configs/config.yaml"

func Init() {
	notify := make(chan bool, 1)
	l := conf.NewLoadConf(notify)
	configInit(l)
	loggerInit()
	dbInit()
	go func() {
		for <-l.Notify {
			global.Config, _ = l.Parse(global.Config)
			loggerInit()
			dbInit()
		}
	}()
}

func configInit(l *conf.LoadConf) {
	err := l.Load(Path)
	if err != nil {
		return
	}
	c := &conf.Conf{}
	c, err = l.Parse(c)
	if err != nil {
		return
	}
	global.Config = c
}

func loggerInit() {
	w := Logger.NewWriter(global.Config.Logger)
	global.Logger = Logger.NewLogger(w)
}

func dbInit() {
	database, err := db.New(global.Config.Db)
	if err != nil {
		log.Print("数据库连接失败")
		return
	}
	global.DB = database
}

func SetLogger(e *echo.Echo) {
	e.Logger = global.Logger
	e.Logger.SetLevel(echoLog.INFO)
}

func SetMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())

	e.Use(middleware.Recover())
}
