package logger

import (
	"strconv"
	"sync"
	"testing"
)

func TestLogger(t *testing.T) {
	conf := getConfig()
	w := NewWriter(conf)
	logger := BuildLogger(conf.TimeFormat, w)
	g := sync.WaitGroup{}
	g.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			logger.Info().Msg("hello world" + strconv.Itoa(i))
			g.Done()
		}(i)
	}
	g.Wait()
}

func getConfig() *Config {
	config := &Config{
		LoggerFilePath: "../../tmp/log",
		LogRotateDate:  1,
		LogRotateSize:  1,
		LogBackupCount: 7,
		Compress:       true,
		TimeFormat:     "2006-01-02 15:04:05",
	}
	return config
}
