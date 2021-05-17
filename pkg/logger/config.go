package logger

// Config Logger config
type Config struct {
	LoggerFilePath string
	LogRotateDate  int
	LogRotateSize  int
	LogBackupCount int
	Compress       bool
	TimeFormat    string
}
