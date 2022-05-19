package log

var loggerDefault *Logger

func init() {
	loggerDefault = NewLoggerDefault()
}

func Warnf(f string, v ...any) {
	loggerDefault.Warnf(f, v...)
}

func Warn(v ...any) {
	loggerDefault.Warn(v...)
}
