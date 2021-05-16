package config

import (
	"testing"
)

const path = "../../configs/config.yaml"

func TestNewViper(t *testing.T) {
	v := NewViper()
	err := v.Load(path)
	if err != nil {
		t.Log(err)
	}
	conf, err := v.Parse(&Config{})
	if err != nil {
		t.Log(err)
	}
	cfg := conf.(*Config)
	t.Log(cfg.Logger.TimeFormat)
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
