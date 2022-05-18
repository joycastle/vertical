package vertical

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"time"
)

type LogConf struct {
	Level uint8
	Fpath string
}

type Logger struct {
	*log.Logger
	level uint8
	fpath string
	Fptr  *os.File
	Fname string
}

const (
	LEVEL_DEBUG uint8 = 1
	LEVEL_INFO  uint8 = 2
	LEVEL_WARN  uint8 = 3
	LEVEL_FATAL uint8 = 4
)

var (
	pid  int = syscall.Getpid()
	ppid int = syscall.Getppid()

	log_prefix_debug = fmt.Sprintf("debug %d ", pid)
	log_prefix_info  = fmt.Sprintf("info %d ", pid)
	log_prefix_warn  = fmt.Sprintf("warn %d ", pid)
	log_prefix_fatal = fmt.Sprintf("fatal %d ", pid)
)

func NewLoggerStdout() *Logger {
	l := &Logger{
		Logger: log.New(os.Stdout, log_prefix_debug, log.Ldate|log.Lmicroseconds|log.Lshortfile),
		Fptr:   os.Stdout,
	}
	return l
}

func NewLogger(level uint8, path string) *Logger {
	var (
		prefix string
		l      *Logger
	)

	switch level {
	case LEVEL_DEBUG:
		prefix = log_prefix_debug
	case LEVEL_INFO:
		prefix = log_prefix_info
	case LEVEL_WARN:
		prefix = log_prefix_warn
	case LEVEL_FATAL:
		prefix = log_prefix_fatal
	default:
		prefix = log_prefix_fatal
	}

	l = &Logger{
		level: level,
		fpath: path,
	}

	l.Logger = log.New(os.Stderr, prefix, log.Ldate|log.Lmicroseconds|log.Lshortfile)
	l.setup_file()

	return l
}

func (l *Logger) setup_file() {
	var (
		fname    string
		fp       *os.File
		deadline time.Time
		err      error
	)

	if len(l.fpath) <= 0 {
		return
	}

	fname, deadline = parse_log_fname(l.fpath)
	if fp, err = open_log_file(fname); err != nil {
		fp = os.Stderr
	}

	l.Fptr = fp
	l.Fname = fname
	l.Logger.SetOutput(fp)

	go func() {
		select {
		case <-time.After(deadline.Sub(time.Now())):
			l.setup_file()
		}
	}()
}

func (l *Logger) Debugf(f string, argv ...interface{}) {
	l.Output(2, fmt.Sprintf(f, argv...))
}

func (l *Logger) Infof(f string, argv ...interface{}) {
	l.Output(2, fmt.Sprintf(f, argv...))
}

func (l *Logger) Warnf(f string, argv ...interface{}) {
	l.Output(2, fmt.Sprintf(f, argv...))
}

func (l *Logger) Fatalf(f string, argv ...interface{}) {
	l.Output(2, fmt.Sprintf(f, argv...))
}
