package config

import (
	"strings"

	cop_config "github.com/joycastle/cop/config"
	cop_conn "github.com/joycastle/cop/connector"
	cop_log "github.com/joycastle/cop/log"
	"github.com/joycastle/vertical/gin"
)

var C_Log map[string]cop_log.LogConf = make(map[string]cop_log.LogConf)
var C_Mysql map[string]cop_conn.MysqlNodeConf = make(map[string]cop_conn.MysqlNodeConf)
var C_Redis map[string]cop_conn.RedisConf = make(map[string]cop_conn.RedisConf)
var C_Gin gin.GinConf

func init() {
	RegisterParser("log.yaml", &C_Log)
	RegisterParser("mysql.yaml", &C_Mysql)
	RegisterParser("redis.yaml", &C_Redis)
	RegisterParser("gin.yaml", &C_Gin)
}

type Parser struct {
	Fname string
	Out   interface{}
}

var fParsers []Parser

func RegisterParser(fname string, out interface{}) {
	fParsers = append(fParsers, Parser{Fname: fname, Out: out})
}

func InitConfig(conf_dir string) error {
	conf_dir = strings.TrimRight(conf_dir, "/")
	for _, parser := range fParsers {
		fileName := conf_dir + "/" + parser.Fname
		if err := cop_config.ReadYmalFromFile(fileName, parser.Out); err != nil && err != cop_config.ErrFileNotExists {
			return err
		}
	}

	return nil
}
