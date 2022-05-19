package vertical

import (
	"testing"
	"time"
)

func init() {
	conf_path := "./conf_file_for_test"
	err := InitConfig(conf_path)
	if err != nil {
		panic(err)
	}
}

func Test_config(t *testing.T) {
	if v, ok := C_Log["error"]; !ok || v.Fpath != "/error.log" || v.Level != 4 {
		t.Logf("load config file failed: %v", C_Log)
		t.Fail()
	}

	if v, ok := C_Mysql["default-master"]; !ok || v.Addr != "127.0.0.1:3306" || v.DnsParams != "charset=utf8mb4&parseTime=True&timeout=1s" || v.Password != "123456" {
		t.Logf("load config file failed: %v", C_Mysql)
		t.Fail()
	}

	if v, ok := C_Mysql["default-slave"]; !ok || v.Addr != "127.1.0.1:13306" || v.DnsParams != "charset=utf8mb4&parseTime=True&timeout=1s" {
		t.Logf("load config file failed: %v", C_Mysql)
		t.Fail()
	}

	if v, ok := C_Redis["main"]; !ok || v.MaxActive != 128 {
		t.Logf("load config file failed: %v", C_Redis)
		t.Fail()
	}

	if C_Gin.ReadTimeout != time.Second*10 || C_Gin.WriteTimeout != time.Second*10 {
		t.Logf("load config file failed: %v", C_Gin)
		t.Fail()
	}
}
