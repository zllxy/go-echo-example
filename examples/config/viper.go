package main

import (
	"fmt"
	"go-vue-admin/pkg/config"
)

const path = "./configs/config.yaml"

func main() {
	v := config.NewViper()
	err := v.Load(path)
	if err != nil {
		fmt.Println(err)
	}
	conf, err := v.Parse(&Config{})
	if err != nil {
		fmt.Println(err)
	}
	cfg := conf.(*Config)
	fmt.Println(cfg.Logger.TimeFormat)
	v.WatchConfig()
	for <-v.Notify {
		fmt.Println(cfg.Logger.LoggerFilePath)
	}
}

type Config struct {
	Logger Logger
}
type Logger struct {
	IsDevMod       bool
	DisableCaller  bool
	DebugLevel     string
	LoggerFilePath string
	LogRotateDate  int
	LogRotateSize  int
	LogBackupCount int
	Compress       bool
	TimeFormat     string
}
