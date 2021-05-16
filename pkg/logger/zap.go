package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

var ZapLog *zap.SugaredLogger

var defaultTimeFormat = "2006-01-02 15:04:05"

var loggerLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"panic": zapcore.PanicLevel,
	"fatal": zapcore.FatalLevel,
}

type ZapLogger struct {
	core         zapcore.Core
	options      []zap.Option
	encoder      zapcore.Encoder
	writer       io.Writer
	writerSyncer zapcore.WriteSyncer
	debugLevel   zapcore.Level
	logger       *zap.Logger
	conf         *Config
}

func InitZap(w io.Writer, conf *Config) {
	if conf.TimeFormat != "" {
		defaultTimeFormat = conf.TimeFormat
	}
	z := &ZapLogger{writer: w, conf: conf}
	z.setEncoder()
	z.setWriterSyncer()
	z.setDebugLevel()
	z.setCore()
	z.setLogger()
	ZapLog = z.logger.Sugar()
}

func (z *ZapLogger) setCore() {
	z.core = zapcore.NewCore(z.encoder, z.writerSyncer, z.debugLevel)
}

func (z *ZapLogger) setEncoder() {
	var encoderConfig zapcore.EncoderConfig

	if z.conf.IsDevMod {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
	}

	encoderConfig.EncodeTime = customTimeEncoder

	encoderConfig.EncodeLevel = customLevelEncoder

	z.encoder = zapcore.NewConsoleEncoder(encoderConfig)
}

func (z *ZapLogger) setWriterSyncer() {
	if z.writer == nil {
		z.writer = os.Stdout
	}
	z.writerSyncer = zapcore.AddSync(z.writer)
}

func (z *ZapLogger) setDebugLevel() {

	level, exist := loggerLevelMap[z.conf.DebugLevel]

	if !exist {
		z.debugLevel = zapcore.DebugLevel
	}

	z.debugLevel = level
}

func (z *ZapLogger) setLogger() {
	z.logger = zap.New(z.core, z.options...)
}

func (z *ZapLogger) setOptions() {

	if !z.conf.DisableCaller {
		z.options = append(z.options, zap.WithCaller(true))
	}

	// 跳过文件调用层数
	z.options = append(z.options, zap.AddCallerSkip(2))
}

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format(defaultTimeFormat) + "]")
}
