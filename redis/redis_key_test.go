package redis

import (
	"testing"
	"time"

	"github.com/joycastle/vertical/connector"
)

func TestCase_RedisKey(t *testing.T) {
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

	key := "joycastle"

	if ret, err := Rds_Del("default", key); err != nil {
		t.Fatal(err, ret)
	}

	if ret, err := Rds_GetString("default", key); err != ErrNotFound {
		t.Fatal(err, ret)
	}

	if ret, err := Rds_Set("default", key, "123456"); err != nil {
		t.Fatal(err, ret)
	}

	if ret, err := Rds_Set("default", key, "123456"); err != nil || ret != "OK" {
		t.Fatal(err, ret)
	}

	if ret, err := Rds_GetString("default", key); err == ErrNotFound {
		t.Fatal(err, ret)
	}

	if ret, err := Rds_Expire("default", key, 2); err != nil || ret != 1 {
		t.Fatal(err, ret)
	}

	time.Sleep(time.Second * 3)

	if ret, err := Rds_GetString("default", key); err != ErrNotFound {
		t.Fatal(err, ret)
	}

	if ret, err := Rds_SetEx("default", key, "123456", 2); err != nil || ret != "OK" {
		t.Fatal(err, ret)
	}

	time.Sleep(3 * time.Second)

	if ret, err := Rds_GetString("default", key); err != ErrNotFound {
		t.Fatal(err, ret)
	}

	if ret, err := Rds_Expire("default", key, 20); err != nil || ret != 0 {
		t.Fatal(err, ret)
	}

	if ret, err := Rds_TTL("default", key+"1"); err != nil || ret != -2 {
		t.Fatal(err, ret)
	}
}
