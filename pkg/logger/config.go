package logger

// Config Logger config
type Config struct {
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
