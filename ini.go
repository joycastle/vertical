package vertical

import (
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/go-ini/ini"
)

func loadConfigFile(dir_path string) (*ini.File, error) {
	dp, err := os.Open(dir_path)
	if err != nil {
		return nil, err
	}

	fnames, err := dp.Readdirnames(0)
	if err != nil {
		return nil, err
	}

	fp := ini.Empty()

	for _, fname := range fnames {
		if !strings.HasSuffix(fname, ".ini") {
			continue
		}

		if content, err := ioutil.ReadFile(dir_path + "/" + fname); err != nil {
			return nil, err
		} else if err := fp.Append(content); err != nil {
			return nil, err
		}
	}

	return fp, nil
}

func GetSection(fp *ini.File, secName string) (*ini.Section, error) {
	return fp.GetSection(secName)
}

func GetSectionValue(sec *ini.Section, key string) (*ini.Key, error) {
	return sec.GetKey(key)
}

func GetSectionValueString(sec *ini.Section, key string) string {
	k, err := GetSectionValue(sec, key)
	if err != nil {
		panic(err)
	}
	return k.MustString("")
}

func GetSectionValueInt(sec *ini.Section, key string) int {
	k, err := GetSectionValue(sec, key)
	if err != nil {
		panic(err)
	}
	return k.MustInt(-1)
}

func GetSectionValueDuration(sec *ini.Section, key string) time.Duration {
	k, err := GetSectionValue(sec, key)
	if err != nil {
		panic(err)
	}
	return k.MustDuration()
}
