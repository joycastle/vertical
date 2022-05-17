package vertical

import (
	"strings"

	"github.com/go-ini/ini"
)

const (
	//log
	CFG_K_LOG       = "Log"
	CFG_K_LOG_LEVEL = "Level"
	CFG_K_LOG_FPATH = "Fpath"

	//mysql
	CFG_K_MYSQL               = "Mysql"
	CFG_K_MYSQL_ADDR          = "Addr"
	CFG_K_MYSQL_USERNAME      = "Username"
	CFG_K_MYSQL_PASSWORD      = "Password"
	CFG_K_MYSQL_DATABASE      = "Database"
	CFG_K_MYSQL_DNSPARAMS     = "DnsParams"
	CFG_K_MYSQL_MAX_IDLE      = "MaxIdle"
	CFG_K_MYSQL_MAX_OPEN      = "MaxOpen"
	CFG_K_MYSQL_MAX_LIFE_TIME = "MaxLifeTime"

	//redis
	CFG_K_REDIS               = "Redis"
	CFG_K_REDIS_ADDRS         = "Addrs"
	CFG_K_REDIS_TEST_INTERVAL = "TestInterval"
	CFG_K_REDIS_MAX_ACTIVE    = "MaxActive"
	CFG_K_REDIS_MAX_IDLE      = "MaxIdle"
	CFG_K_REDIS_IDLE_TIMEOUT  = "IdleTimeout"
	CFG_K_REDIS_CONN_TIMEOUT  = "ConnectTimeout"
	CFG_K_REDIS_READ_TIMEOUT  = "ReadTimeout"
	CFG_K_REDIS_WRITE_TIMEOUT = "WriteTimeout"

	//gin
	CFG_K_GIN               = "Gin"
	CFG_K_GIN_READ_TIMEOUT  = "ReadTimeout"
	CFG_K_GIN_WRITE_TIMEOUT = "WriteTimeout"
)

//global vars
var C_Log map[string]LogConf = make(map[string]LogConf)
var C_Mysql map[string]MysqlConf = make(map[string]MysqlConf)
var C_Redis map[string]RedisConf = make(map[string]RedisConf)
var C_Gin GinConf

func init() {
	RegisterParseMethod(parser_log)
	RegisterParseMethod(parser_mysql)
	RegisterParseMethod(parser_redis)
	RegisterParseMethod(parser_gin)
}

func parser_log(fp *ini.File) error {
	psec, err := GetSection(fp, CFG_K_LOG)
	if err != nil {
		return ErrSectionNotExists
	}

	for _, sec := range psec.ChildSections() {
		sn := strings.Replace(sec.Name(), CFG_K_LOG+".", "", -1)

		c := LogConf{
			Level: uint8(GetSectionValueInt(sec, CFG_K_LOG_LEVEL)),
			Fpath: GetSectionValueString(sec, CFG_K_LOG_FPATH),
		}

		C_Log[sn] = c
	}

	return nil
}

func parser_mysql(fp *ini.File) error {
	psec, err := GetSection(fp, CFG_K_MYSQL)
	if err != nil {
		return ErrSectionNotExists
	}

	for _, sec := range psec.ChildSections() {
		sn := strings.Replace(sec.Name(), CFG_K_MYSQL+".", "", -1)

		c := MysqlConf{
			Addr:        GetSectionValueString(sec, CFG_K_MYSQL_ADDR),
			Username:    GetSectionValueString(sec, CFG_K_MYSQL_USERNAME),
			Password:    GetSectionValueString(sec, CFG_K_MYSQL_PASSWORD),
			Database:    GetSectionValueString(sec, CFG_K_MYSQL_DATABASE),
			DnsParams:   GetSectionValueString(sec, CFG_K_MYSQL_DNSPARAMS),
			MaxIdle:     GetSectionValueInt(sec, CFG_K_MYSQL_MAX_IDLE),
			MaxOpen:     GetSectionValueInt(sec, CFG_K_MYSQL_MAX_OPEN),
			MaxLifeTime: GetSectionValueDuration(sec, CFG_K_MYSQL_MAX_LIFE_TIME),
		}

		C_Mysql[sn] = c
	}

	return nil
}

func parser_redis(fp *ini.File) error {
	psec, err := GetSection(fp, CFG_K_REDIS)
	if err != nil {
		return ErrSectionNotExists
	}

	for _, sec := range psec.ChildSections() {
		sn := strings.Replace(sec.Name(), CFG_K_REDIS+".", "", -1)

		c := RedisConf{
			Addrs:          strings.Split(GetSectionValueString(sec, CFG_K_REDIS_ADDRS), ","),
			TestInterval:   GetSectionValueDuration(psec, CFG_K_REDIS_TEST_INTERVAL),
			MaxActive:      GetSectionValueInt(psec, CFG_K_REDIS_MAX_ACTIVE),
			MaxIdle:        GetSectionValueInt(psec, CFG_K_REDIS_MAX_IDLE),
			IdleTimeout:    GetSectionValueDuration(psec, CFG_K_REDIS_IDLE_TIMEOUT),
			ConnectTimeout: GetSectionValueDuration(psec, CFG_K_REDIS_CONN_TIMEOUT),
			ReadTimeout:    GetSectionValueDuration(psec, CFG_K_REDIS_READ_TIMEOUT),
			WriteTimeout:   GetSectionValueDuration(psec, CFG_K_REDIS_WRITE_TIMEOUT),
		}

		C_Redis[sn] = c
	}

	return nil
}

func parser_gin(fp *ini.File) error {
	sec, err := GetSection(fp, CFG_K_GIN)
	if err != nil {
		return ErrSectionNotExists
	}

	C_Gin = GinConf{
		ReadTimeout:  GetSectionValueDuration(sec, CFG_K_GIN_READ_TIMEOUT),
		WriteTimeout: GetSectionValueDuration(sec, CFG_K_GIN_WRITE_TIMEOUT),
	}

	return nil
}
