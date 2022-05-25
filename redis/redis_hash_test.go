package redis

import (
	"testing"
	"time"

	"github.com/joycastle/cop/connector"
)

func TestCase_RedisHash(t *testing.T) {
	configs := make(map[string]connector.RedisConf)

	configs["default"] = connector.RedisConf{
		Addr:           "127.0.0.1:6379,127.0.0.1:6379",
		Password:       "123456",
		MaxActive:      64,
		MaxIdle:        16,
		IdleTimeout:    time.Second * 240,
		ConnectTimeout: time.Second * 10,
		ReadTimeout:    time.Second,
		WriteTimeout:   time.Second,
		TestInterval:   time.Second * 60,
	}

	connector.InitRedisConn(configs)

	key := "joycastle_hash"

	if ret, err := Rds_Del("default", key); err != nil {
		t.Fatal(err, ret)
	}

	if ret, err := Rds_HGetString("default", key, "levin"); err != ErrNotFound {
		t.Fatal(ret, err)
	}

	if ret, err := Rds_HSet("default", key, "levin", "123"); err != nil || ret != 1 {
		t.Fatal(ret, err)
	}

	if ret, err := Rds_HExists("default", key, "levin"); err != nil || ret != 1 {
		t.Fatal(ret, err)
	}

	if ret, err := Rds_HDel("default", key, "levin"); err != nil {
		t.Fatal(ret, err)
	}

	if ret, err := Rds_HExists("default", key, "levin"); err != nil || ret != 0 {
		t.Fatal(ret, err)
	}

	if ret, err := Rds_HGetAllString("default", key); err != nil || len(ret) != 0 {
		t.Fatal(ret, err)
	}

	if ret, err := Rds_HGetAllString("default", key+"1"); err != nil {
		t.Fatal(ret, err)
	}

	if ret, err := Rds_HSet("default", key, "levin", "123"); err != nil || ret != 1 {
		t.Fatal(ret, err)
	}

	if ret, err := Rds_HSet("default", key, "blue", "000"); err != nil || ret != 1 {
		t.Fatal(ret, err)
	}

	if ret, err := Rds_HSet("default", key, "green", "111"); err != nil || ret != 1 {
		t.Fatal(ret, err)
	}

	if ret, err := Rds_HGetAllString("default", key); err != nil || ret["green"] != "111" {
		t.Fatal(ret, err)
	}
}
