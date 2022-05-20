package config

import (
	"testing"
	"time"
)

func TestCase_App_Ymal(t *testing.T) {
	InitYmalConfig("./")

	if C_Log["error"] != "./error.log" {
		t.Fatal(C_Log)
	}

	if C_Mysql["db_game"].Master.Dsn != "root:123456@tcp(127.0.0.1:3306)/db_game?charset=utf8mb4&parseTime=True&timeOut=10s" {
		t.Fatal(C_Mysql)
	}

	if C_Redis["default"].TestInterval != 60*time.Second || C_Redis["default"].ConnectTimeout != time.Second*10 {
		t.Fatal(C_Redis)
	}

	if C_Gin.WriteTimeout != time.Second*30 {
		t.Fatal(C_Redis)
	}
}
