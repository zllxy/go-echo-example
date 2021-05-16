package logger

import (
	"github.com/natefinch/lumberjack"
	"strconv"
	"time"
)

type Writer struct {
	Logger   *lumberjack.Logger
	FileName func() string
}

func NewWriter(cfg *Config) *Writer {
	logger := &lumberjack.Logger{
		MaxSize:    cfg.LogRotateSize,
		MaxBackups: cfg.LogBackupCount,
		MaxAge:     cfg.LogRotateDate,
		Compress:   cfg.Compress,
	}
	return &Writer{
		Logger:   logger,
		FileName: getFileName(cfg),
	}
}
func getFileName(cfg *Config) func() string {
	return func() string {
		timeNow := time.Now()
		fileDir := timeNow.Format("200601")
		return cfg.LoggerFilePath + "/" + fileDir + "/" + strconv.Itoa(timeNow.Day()) + ".log"
	}
}

func (l *Writer) Write(p []byte) (n int, err error) {
	l.Logger.Filename = l.FileName()
	n, err = l.Logger.Write(p)
	return n, err
}