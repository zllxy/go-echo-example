package log

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
	"io"
	"os"
)

var DefaultTimeFormat = "2006-01-02 15:04:05"

type Logger struct {
	*log.Logger
	ZeroLog zerolog.Logger
}

func NewLogger(w io.Writer) *Logger {
	l := &Logger{
		Logger:  log.New("-"),
		ZeroLog: buildZeroLog(w),
	}
	return l
}

func buildZeroLog(w io.Writer) zerolog.Logger {
	zerolog.TimeFieldFormat = DefaultTimeFormat
	var writer zerolog.LevelWriter
	if w != nil {
		writer = zerolog.MultiLevelWriter(w)
	} else {
		writer = zerolog.MultiLevelWriter(os.Stdout)
	}

	return zerolog.New(writer).With().Caller().Timestamp().Logger()
}

func (l *Logger) SetOutput(writer io.Writer) {
	l.Logger.SetOutput(writer)
	l.ZeroLog.Output(writer)
}

func (l *Logger) SetLevel(level log.Lvl) {
	l.Logger.SetLevel(level)
	if level == log.OFF {
		l.ZeroLog = l.ZeroLog.Level(zerolog.Disabled)
	} else {
		zeroLevel := int8(level) - 1
		l.ZeroLog = l.ZeroLog.Level(zerolog.Level(zeroLevel))
	}
}

func (l *Logger) Info(i ...interface{}) {
	l.ZeroLog.Info().Msg(fmt.Sprint(i...))
}

func (l *Logger) Debug(i ...interface{}) {
	l.ZeroLog.Debug().Msg(fmt.Sprint(i...))
}

func (l *Logger) Warn(i ...interface{}) {
	l.ZeroLog.Warn().Msg(fmt.Sprint(i...))
}

func (l *Logger) Error(i ...interface{}) {
	l.ZeroLog.Error().Msg(fmt.Sprint(i...))
}

func (l *Logger) Fatal(i ...interface{}) {
	l.ZeroLog.Fatal().Msg(fmt.Sprint(i...))
}

func (l *Logger) Print(i ...interface{}) {
	l.ZeroLog.Print(i...)
}
