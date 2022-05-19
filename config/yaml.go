package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	log "github.com/joycastle/vertical/logger"
	"github.com/joycastle/vertical/util"
	"gopkg.in/yaml.v2"
)

type Parser struct {
	Fname string
	Out   interface{}
}

var fParsers []Parser

func RegisterParser(fname string, out interface{}) {
	fParsers = append(fParsers, Parser{Fname: fname, Out: out})
}

func ReadYmalFromFile(fileName string, out interface{}) error {
	fd, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	if out == nil {
		return fmt.Errorf("yaml 解析失败：输出文件是空的!!")
	}

	if err := yaml.Unmarshal(fd, out); err != nil {
		return fmt.Errorf("yaml 解析失败：%s,  ", err.Error())
	}
	return nil
}

func InitYmalConfig(conf_dir string) {
	conf_dir = filepath.Dir(conf_dir)
	for _, parser := range fParsers {

		fileName := conf_dir + "/" + parser.Fname
		if !util.FileExists(fileName) {
			log.Warnf("file not exists: %s", fileName)
			continue
		}

		if !strings.HasSuffix(fileName, ".ymal") && !strings.HasSuffix(fileName, ".yml") {
			log.Warnf("The file format is incorrect must be .ymal or .yml: %s", fileName)
			continue
		}

		if err := ReadYmalFromFile(fileName, parser.Out); err != nil {
			panic(err)
		}
	}
}
