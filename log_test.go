package log

import (
	"testing"
)

var logger *Logger
var logger_f *Logger

func init() {
	logger = NewLoggerStdout()
	logger_f = NewLogger(2, "./test.log")
}

func TestCase_log_file_rotate(t *testing.T) {
	for i := 0; i < 10; i++ {
		logger.Printf("log test %d", i)
	}
}

func Benchmark_Debugf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logger_f.Printf("Hi: %s", "Jack")
	}
}
