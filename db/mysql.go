package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

type MysqlConfig struct {
	Type        string `yaml:"type"`
	Master      string `yaml:"master"`
	TablePrefix string `yaml:"table-prefix"`
	MaxIdleConn int    `yaml:"max-idle-conn"`
	MaxOpenConn int    `yaml:"max-open-conn"`
	LogMode     bool   `yaml:"log-mode"`
}

func InitMysql(conf *MysqlConfig) (err error) {
	DB, err = gorm.Open(conf.Type, conf.Master)
	if err != nil {
		return
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.TablePrefix + defaultTableName
	}

	DB.LogMode(conf.LogMode)
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	return
}
