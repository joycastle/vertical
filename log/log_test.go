package log

import (
	"errors"
	"testing"
)

func Test_log(t *testing.T) {
	// t.Fatal("not implemented")

	log := NewLoggerDefault()
	log.Printf("This is Logger:%s", "hello")
	log.Println("This is Logger", errors.New("Hello"))
	log.Infof("This is Logger:%s", "hello")
	log.Info("This is Logger", errors.New("Hello"))
	log.Debugf("This is Logger:%s", "hello")
	log.Debug("This is Logger", errors.New("Hello"))
	log.Warnf("This is Logger:%s", "hello")
	log.Warn("This is Logger", errors.New("Hello"))
	log.Fatalf("This is Logger:%s", "hello")
	log.Fatal("This is Logger", errors.New("Hello"))

	DisableColor()

	log.Printf("This is Logger:%s", "hello")
	log.Println("This is Logger", errors.New("Hello"))
	log.Infof("This is Logger:%s", "hello")
	log.Info("This is Logger", errors.New("Hello"))
	log.Debugf("This is Logger:%s", "hello")
	log.Debug("This is Logger", errors.New("Hello"))
	log.Warnf("This is Logger:%s", "hello")
	log.Warn("This is Logger", errors.New("Hello"))
	log.Fatalf("This is Logger:%s", "hello")
	log.Fatal("This is Logger", errors.New("Hello"))

	EnableColor()
}

func Benchmark_log(b *testing.B) {
	//DisableColor()
	log := NewLoggerDefault()
	for n := 0; n < b.N; n++ {
		log.Fatalf("This is Logger:%s", "hello")
	}
}
