package config

import (
	"github.com/joycastle/vertical/connector"
	"github.com/joycastle/vertical/gin"
)

var C_Log map[string]string = make(map[string]string)
var C_Mysql map[string]connector.MysqlNodeConf = make(map[string]connector.MysqlNodeConf)
var C_Redis map[string]connector.RedisConf = make(map[string]connector.RedisConf)
var C_Gin gin.GinConf

func init() {
	RegisterParser("log.yaml", &C_Log)
	RegisterParser("mysql.yaml", &C_Mysql)
	RegisterParser("redis.yaml", &C_Redis)
	RegisterParser("gin.yaml", &C_Gin)
}
