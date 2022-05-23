package log

import (
	"github.com/joycastle/cop/log"
)

var (
	loggerMapping map[string]*log.Logger = make(map[string]*log.Logger)
)

func InitLogger(configs map[string]log.LogConf) {
	for sn, logConf := range configs {
		loggerMapping[sn] = log.NewLogger(logConf)
	}
}

func GetLogger(sn string) *log.Logger {
	if v, ok := loggerMapping[sn]; ok {
		return v
	}
	return log.Default
}
