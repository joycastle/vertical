package config

import "github.com/joycastle/vertical/connector"

var C_Log map[string]string = make(map[string]string)
var C_Mysql map[string]connector.MysqlNodeConf = make(map[string]connector.MysqlNodeConf)
var C_Redis map[string]connector.RedisConf = make(map[string]connector.RedisConf)

func init() {
	RegisterParser("log.ymal", &C_Log)
	RegisterParser("mysql.ymal", &C_Mysql)
	RegisterParser("redis.ymal", &C_Redis)
}
