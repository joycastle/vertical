package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

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
		if err := ReadYmalFromFile(conf_dir+"/"+parser.Fname, parser.Out); err != nil {
			panic(err)
		}
	}
}
