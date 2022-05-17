package component

import (
	"testing"
	"vertical/config"
)

func Test_Log_config(t *testing.T) {
	conf_path := "./conf_file_for_test"
	err := config.InitConfig(conf_path)
	if err != nil {
		t.Logf("load config file failed: %s", err)
		t.Fail()
	}

	//map[debug:{3 /debug.log-*-*-*} error:{12 /error.log-*-*-*} normal:{3 /main.log-*-*-*} wf:{12 /main.wf.log-*-*-*}]
	if v, ok := config.C_Log["debug"]; !ok || v.Fpath != "/debug.log-*-*-*" || v.Level != 3 {
		t.Logf("load config file failed: %v", config.C_Log)
		t.Fail()
	}
}
