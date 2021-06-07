package log

import (
	"strconv"
	"sync"
	"testing"
)

func TestLogger(t *testing.T) {
	conf := getConfig()
	w := NewWriter(conf)
	logger := NewLogger(w)
	g := sync.WaitGroup{}
	g.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			logger.ZeroLog.Info().Msg("hello world" + strconv.Itoa(i))
			logger.Info("hello log" + strconv.Itoa(i))
			g.Done()
		}(i)
	}
	g.Wait()
}

func getConfig() *Conf {
	config := &Conf{
		LoggerFilePath: "../../tmp/log",
		LogRotateDate:  1,
		LogRotateSize:  1,
		LogBackupCount: 7,
		Compress:       true,
	}
	return config
}
