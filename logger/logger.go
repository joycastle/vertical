package logger

import "github.com/joycastle/vertical/log"

var (
	loggerDefault *log.Logger
	loggerMapping map[string]*log.Logger = make(map[string]*log.Logger)
)

func init() {
	loggerDefault = log.NewLoggerDefault()
}

func InitLogger(configs map[string]string) {
	for sn, logPath := range configs {
		loggerMapping[sn] = log.NewLogger(logPath)
	}
}

func GetLogger(sn string) *log.Logger {
	if v, ok := loggerMapping[sn]; ok {
		return v
	}
	return loggerDefault
}

func Warnf(f string, v ...any) {
	loggerDefault.Warnf(f, v...)
}

func Warn(v ...any) {
	loggerDefault.Warn(v...)
}
