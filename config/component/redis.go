package component

import (
	"strings"
	"vertical/config"
	"vertical/redis"

	"github.com/go-ini/ini"
)

const (
	CFG_K_REDIS               = "Redis"
	CFG_K_REDIS_ADDRS         = "Addrs"
	CFG_K_REDIS_TEST_INTERVAL = "TestInterval"
	CFG_K_REDIS_MAX_ACTIVE    = "MaxActive"
	CFG_K_REDIS_MAX_IDLE      = "MaxIdle"
	CFG_K_REDIS_IDLE_TIMEOUT  = "IdleTimeout"
	CFG_K_REDIS_CONN_TIMEOUT  = "ConnectTimeout"
	CFG_K_REDIS_READ_TIMEOUT  = "ReadTimeout"
	CFG_K_REDIS_WRITE_TIMEOUT = "WriteTimeout"
)

func init() {
	config.RegisterParseMethod(parser_redis)
}

func parser_redis(fp *ini.File) error {
	psec, err := config.GetSection(fp, CFG_K_REDIS)
	if err != nil {
		return config.ErrSectionNotExists
	}

	for _, sec := range psec.ChildSections() {
		sn := strings.Replace(sec.Name(), CFG_K_REDIS+".", "", -1)

		c := redis.RedisConf{
			Addrs:          strings.Split(config.GetSectionValueString(sec, CFG_K_REDIS_ADDRS), ","),
			TestInterval:   config.GetSectionValueDuration(psec, CFG_K_REDIS_TEST_INTERVAL),
			MaxActive:      config.GetSectionValueInt(psec, CFG_K_REDIS_MAX_ACTIVE),
			MaxIdle:        config.GetSectionValueInt(psec, CFG_K_REDIS_MAX_IDLE),
			IdleTimeout:    config.GetSectionValueDuration(psec, CFG_K_REDIS_IDLE_TIMEOUT),
			ConnectTimeout: config.GetSectionValueDuration(psec, CFG_K_REDIS_CONN_TIMEOUT),
			ReadTimeout:    config.GetSectionValueDuration(psec, CFG_K_REDIS_READ_TIMEOUT),
			WriteTimeout:   config.GetSectionValueDuration(psec, CFG_K_REDIS_WRITE_TIMEOUT),
		}

		config.C_Redis[sn] = c
	}

	return nil
}
