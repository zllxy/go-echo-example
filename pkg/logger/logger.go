package logger

import (
	"github.com/rs/zerolog"
	"io"
	"os"
)

var defaultTimeFormat = "2006-01-02 15:04:05"

type Logger struct {
	writer  io.Writer
	ZeroLog zerolog.Logger
}

func BuildLogger(timeFormat string, w io.Writer) *zerolog.Logger {
	zerolog.TimeFieldFormat = defaultTimeFormat
	if timeFormat != "" {
		zerolog.TimeFieldFormat = timeFormat
	}
	l := &Logger{}
	l.setWriter(w)
	l.setZeroLog()
	return &l.ZeroLog
}

func (l *Logger) setWriter(w io.Writer) {
	var writers []io.Writer
	writers = append(writers, os.Stdout)
	if w != nil {
		writers = append(writers, w)
	}

	l.writer = zerolog.MultiLevelWriter(writers...)
}

func (l *Logger) setZeroLog() {
	l.ZeroLog = zerolog.New(l.writer).With().Caller().Timestamp().Logger()
}
