package vertical

import (
	"testing"
	"time"
)

func init() {
	configs := make(map[string]RedisConf)
	configs["default"] = RedisConf{
		Addrs: []string{"127.0.0.1:6379"},

		Password:     "123456",
		TestInterval: time.Second * 60,

		MaxActive:   128,
		MaxIdle:     16,
		IdleTimeout: time.Second * 240,

		ConnectTimeout: time.Second * 10,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
	}

	err := InitRedis(configs)
	if err != nil {
		panic(err)
	}
}

func Test_Key(t *testing.T) {
	sn, err := GetRedis("default")
	if err != nil {
		t.Logf("get redis error:%s", err)
		t.Fail()
	}

	if ret, err := sn.Set("joycastle", "123456"); err != nil || ret != "OK" {
		t.Logf("set error: %s", err)
		t.Fail()
	}

	if ret, err := sn.Del("joycastle"); err != nil || ret == false {
		t.Logf("DEL error: %s", err)
		t.Fail()
	}

	if ret, err := sn.Del("joycastle"); err != nil || ret == true {
		t.Logf("DEL error: %s", err)
		t.Fail()
	}

	if ret, err := sn.Set("joycastle", "123456"); err != nil || ret != "OK" {
		t.Logf("set error: %s", err)
		t.Fail()
	}

	if ret, err := sn.GetString("joycastle"); err != nil || ret != "123456" {
		t.Logf("GetString error: %s", err)
		t.Fail()
	}

	if ret, err := sn.GetInt("joycastle"); err != nil || ret != 123456 {
		t.Logf("GetInt error: %s", err)
		t.Fail()
	}

	if ret, err := sn.Expire("joycastle", 1000); err != nil || ret == false {
		t.Logf("Expire error: %s", err)
		t.Fail()
	}

	if ret, err := sn.TTL("joycastle"); err != nil || ret <= 2 {
		t.Logf("TTL error: %s", err)
		t.Fail()
	}

	if ret, err := sn.SetEx("joycastle", "123888", 86400); err != nil || ret != "OK" {
		t.Logf("SetEx error: %s", err)
		t.Fail()
	}
}
