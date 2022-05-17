package log

var (
	stubLogger *Logger            = NewLoggerStdout()
	loggers    map[string]*Logger = make(map[string]*Logger)
)

func InitLogs(configs map[string]LogConf) {
	for sn, config := range configs {
		loggers[sn] = NewLogger(config.Level, config.Fpath)
	}
}

func GetLogger(sn string) *Logger {
	if logger, ok := loggers[sn]; ok {
		return logger
	}
	return stubLogger
}
