package vertical

import (
	"testing"
)

var t_logger *Logger
var t_logger_f *Logger

func init() {
	t_logger = NewLoggerStdout()
	t_logger_f = NewLogger(2, "./test.log")
}

func TestCase_log_file_rotate(t *testing.T) {
	for i := 0; i < 10; i++ {
		t_logger.Printf("log test %d", i)
	}
}

func Benchmark_Debugf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t_logger_f.Printf("Hi: %s", "Jack")
	}
}
