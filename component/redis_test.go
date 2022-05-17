package component

import (
	"fmt"
	"testing"
	"vertical/config"
)

func Test_Redis_config(t *testing.T) {
	conf_path := "./conf_file_for_test"
	err := config.InitConfig(conf_path)
	if err != nil {
		t.Logf("load config file failed: %s", err)
		t.Fail()
	}

	if v, ok := config.C_Redis["main"]; !ok || v.MaxActive != 128 {
		t.Logf("load config file failed: %v", config.C_Redis)
		t.Fail()
	}

	fmt.Println(config.C_Redis)
}
