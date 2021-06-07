package log

// Conf Logger conf
type Conf struct {
	LoggerFilePath string
	LogRotateDate  int
	LogRotateSize  int
	LogBackupCount int
	Compress       bool
}
