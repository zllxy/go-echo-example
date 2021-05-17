package logger

// Config Logger config
type Config struct {
	LogConf    *LogConf
	WriterConf *WriterConf
}

type LogConf struct {
	IsDevMod      bool
	DisableCaller bool
	DebugLevel    string
	TimeFormat    string
}

type WriterConf struct {
	LoggerFilePath string
	LogRotateDate  int
	LogRotateSize  int
	LogBackupCount int
	Compress       bool
}
