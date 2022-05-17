package config

import (
	"vertical/log"
	"vertical/orm"
	"vertical/redis"
)

var C_Log map[string]log.LogConf = make(map[string]log.LogConf)
var C_Mysql map[string]orm.MysqlConf = make(map[string]orm.MysqlConf)
var C_Redis map[string]redis.RedisConf = make(map[string]redis.RedisConf)
