package config

import (
	"testing"
	"time"
)

func Test_loadConfigFile(t *testing.T) {
	conf_path := "./component/conf_file_for_test"
	fp, err := loadConfigFile(conf_path)
	if err != nil {
		t.Logf("load config file failed: %s", err)
		t.Fail()
	}

	sec, err := GetSection(fp, "Demo")
	if err != nil {
		t.Logf("section get faild: %s", err)
	}

	vStr := GetSectionValueString(sec, "Name")
	t.Logf("Name=%s", vStr)
	if vStr != "config_for_test" {
		t.Fail()
	}
	vInt := GetSectionValueInt(sec, "MaxIdel")
	t.Logf("MaxIdel=%d", vInt)
	if vInt != 199 {
		t.Fail()
	}
	vTimeDuration := GetSectionValueDuration(sec, "TimeOut")
	t.Logf("TimeOut=%v", vTimeDuration)
	if vTimeDuration != time.Duration(300*time.Second) {
		t.Fail()
	}

}
