package component

import (
	"strings"
	"vertical/config"
	"vertical/log"

	"github.com/go-ini/ini"
)

const (
	CFG_K_LOG       = "Log"
	CFG_K_LOG_LEVEL = "Level"
	CFG_K_LOG_FPATH = "Fpath"
)

func init() {
	config.RegisterParseMethod(parser_log)
}

func parser_log(fp *ini.File) error {
	psec, err := config.GetSection(fp, CFG_K_LOG)
	if err != nil {
		return config.ErrSectionNotExists
	}

	for _, sec := range psec.ChildSections() {
		sn := strings.Replace(sec.Name(), CFG_K_LOG+".", "", -1)

		c := log.LogConf{
			Level: uint8(config.GetSectionValueInt(sec, CFG_K_LOG_LEVEL)),
			Fpath: config.GetSectionValueString(sec, CFG_K_LOG_FPATH),
		}

		config.C_Log[sn] = c
	}

	return nil
}
