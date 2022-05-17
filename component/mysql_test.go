package component

import (
	"testing"
	"vertical/config"
)

func Test_Mysql_config(t *testing.T) {
	conf_path := "./conf_file_for_test"
	err := config.InitConfig(conf_path)
	if err != nil {
		t.Logf("load config file failed: %s", err)
		t.Fail()
	}

	if v, ok := config.C_Mysql["default-master"]; !ok || v.Addr != "127.0.0.1:3306" || v.DnsParams != "charset=utf8mb4&parseTime=True" {
		t.Logf("load config file failed: %v", config.C_Mysql)
		t.Fail()
	}
}
