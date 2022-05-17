package component

import (
	"strings"
	"vertical/config"
	"vertical/orm"

	"github.com/go-ini/ini"
)

const (
	CFG_K_MYSQL               = "Mysql"
	CFG_K_MYSQL_ADDR          = "Addr"
	CFG_K_MYSQL_USERNAME      = "Username"
	CFG_K_MYSQL_PASSWORD      = "Password"
	CFG_K_MYSQL_DATABASE      = "Database"
	CFG_K_MYSQL_DNSPARAMS     = "DnsParams"
	CFG_K_MYSQL_MAX_IDLE      = "MaxIdle"
	CFG_K_MYSQL_MAX_OPEN      = "MaxOpen"
	CFG_K_MYSQL_MAX_LIFE_TIME = "MaxLifeTime"
)

func init() {
	config.RegisterParseMethod(parser_mysql)
}

func parser_mysql(fp *ini.File) error {
	psec, err := config.GetSection(fp, CFG_K_MYSQL)
	if err != nil {
		return config.ErrSectionNotExists
	}

	for _, sec := range psec.ChildSections() {
		sn := strings.Replace(sec.Name(), CFG_K_MYSQL+".", "", -1)

		c := orm.MysqlConf{
			Addr:        config.GetSectionValueString(sec, CFG_K_MYSQL_ADDR),
			Username:    config.GetSectionValueString(sec, CFG_K_MYSQL_USERNAME),
			Password:    config.GetSectionValueString(sec, CFG_K_MYSQL_PASSWORD),
			Database:    config.GetSectionValueString(sec, CFG_K_MYSQL_DATABASE),
			DnsParams:   config.GetSectionValueString(sec, CFG_K_MYSQL_DNSPARAMS),
			MaxIdle:     config.GetSectionValueInt(sec, CFG_K_MYSQL_MAX_IDLE),
			MaxOpen:     config.GetSectionValueInt(sec, CFG_K_MYSQL_MAX_OPEN),
			MaxLifeTime: config.GetSectionValueDuration(sec, CFG_K_MYSQL_MAX_LIFE_TIME),
		}

		config.C_Mysql[sn] = c
	}

	return nil
}
