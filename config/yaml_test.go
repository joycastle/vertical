package config

import (
	"fmt"
	"testing"
)

type MysqlConf struct {
	User     string `yaml:"user"`
	PassWord string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int32  `yaml:"port"`
	DbName   string `yaml:"dbname"`
}

type Config struct {
	Mysql MysqlConf `yaml:"mysql"`
}

/*
func TestCase_Parse_Ymal(t *testing.T) {
	var mc Config
	RegisterParser(&mc)
	InitYmalConfig("./test.ymal")
	if mc.Mysql.User != "root" || mc.Mysql.PassWord != "mypassword" || mc.Mysql.Port != 3306 {
		t.Fail()
	}
}*/

func TestCase_App_Ymal(t *testing.T) {
	InitYmalConfig("./")
	fmt.Println(C_Log)
}
