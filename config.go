package config

import (
	"errors"

	"github.com/go-ini/ini"
)

var (
	parseMethods []func(*ini.File) error = []func(*ini.File) error{}
)

var (
	ErrSectionNotExists = errors.New("section not exists")
)

func RegisterParseMethod(f func(*ini.File) error) {
	parseMethods = append(parseMethods, f)
}

func InitConfig(dir_path string) error {
	iFile, err := loadConfigFile(dir_path)
	if err != nil {
		return err
	}

	for _, parseMethod := range parseMethods {
		if parseMethod != nil {
			if err := parseMethod(iFile); err != nil && err != ErrSectionNotExists {
				return err
			}
		}
	}

	return nil
}
