package logger

import (
	"strconv"
	"sync"
	"testing"
)

func TestZap(t *testing.T) {
	conf := getConfig()
	w := NewWriter(conf.WriterConf)
	zap := NewZapLogger(w, conf.LogConf)
	sugar := zap.logger.Sugar()
	g := sync.WaitGroup{}
	g.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			sugar.Info("hello world" + strconv.Itoa(i))
			g.Done()
		}(i)
	}
	g.Wait()
}

func getConfig() *Config {
	config := &Config{}
	config.WriterConf = &WriterConf{
		LoggerFilePath: "../../tmp/log",
		LogRotateDate:  1,
		LogRotateSize:  1,
		LogBackupCount: 7,
		Compress:       true,
	}
	config.LogConf = &LogConf{
		IsDevMod:      true,
		DisableCaller: false,
		DebugLevel:    "debug",
		TimeFormat:    "2006-01-02 15:04:05",
	}
	return config
}
