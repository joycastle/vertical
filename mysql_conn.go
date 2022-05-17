package vertical

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"

	"gorm.io/gorm/logger"
)

const (
	LOG_MODE_DEBUG       = "debug"
	LOG_MODE_PROD        = "prod"
	CFG_Node_Type_Master = "master"
	CFG_Node_Type_Slave  = "slave"
)

var (
	modeType         string              = LOG_MODE_PROD
	mysqlConnMapping map[string]*gorm.DB = make(map[string]*gorm.DB)
)

type MysqlConf struct {
	Addr      string
	Username  string
	Password  string
	Database  string
	DnsParams string

	MaxIdle     int
	MaxOpen     int
	MaxLifeTime time.Duration
}

func LogMode(m string) {
	modeType = m
}

func InitMysql(configs map[string]MysqlConf) error {

	var rwConfMapping map[string]map[string]MysqlConf = make(map[string]map[string]MysqlConf)

	for sn, config := range configs {
		snv := strings.Split(sn, "-")
		if len(snv) != 2 {
			continue
		}
		node := snv[0]
		nType := snv[1]
		if _, ok := rwConfMapping[node]; !ok {
			rwConfMapping[node] = make(map[string]MysqlConf)
		}
		rwConfMapping[node][nType] = config
	}

	for node, cfgs := range rwConfMapping {

		masterConfig := cfgs["master"]
		slaveConfig := cfgs["slave"]

		masterDsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", masterConfig.Username, masterConfig.Password, masterConfig.Addr, masterConfig.Database, masterConfig.DnsParams)
		slaveDsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", slaveConfig.Username, slaveConfig.Password, slaveConfig.Addr, slaveConfig.Database, slaveConfig.DnsParams)

		masterDsnConfig := mysql.Config{
			DSN:                       masterDsn,
			DefaultStringSize:         64,    // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		}

		slaveDsnConfig := mysql.Config{
			DSN:                       slaveDsn,
			DefaultStringSize:         64,    // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		}

		gormConfig := &gorm.Config{
			Logger: logger.New(
				GetLogger("slow"),
				logger.Config{
					//SlowThreshold:             100 * time.Millisecond, // Slow SQL threshold
					LogLevel:                  logger.Info, // Log level
					IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
					Colorful:                  false,       // Disable color
				},
			),
		}

		db, err := gorm.Open(mysql.New(masterDsnConfig), gormConfig)
		if err != nil {
			return fmt.Errorf("mysql init error %s", masterDsn)
		}

		db.Use(dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.New(masterDsnConfig)},
			Replicas: []gorm.Dialector{mysql.New(slaveDsnConfig)},
			Policy:   dbresolver.RandomPolicy{}, // sources/replicas 负载均衡策略
		}).SetConnMaxLifetime(masterConfig.MaxLifeTime).
			SetMaxIdleConns(masterConfig.MaxIdle).
			SetMaxOpenConns(masterConfig.MaxOpen))

		mysqlConnMapping[node] = db
	}

	return nil
}

func GetMysql(node string) *gorm.DB {
	return mysqlConnMapping[node]
}
