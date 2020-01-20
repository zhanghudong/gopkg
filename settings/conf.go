package settings

import (
	"github.com/zhanghudong/gopkg/db"
	"github.com/zhanghudong/gopkg/sms/aliyun"
	"github.com/zhanghudong/gopkg/storage/oss"
	"time"
)

/**
 * @Author: zhanghudong
 * @Date: 2019-12-13 18:05
 */

var ApplicationConfig *Config

type Config struct {
	App    App               `yaml:"app"`
	Server Server            `yaml:"server"`
	Oss    *oss.Config       `yaml:"oss"`
	Mysql  *db.MysqlConfig   `yaml:"mysql"`
	Sms    *aliyun.SmsConfig `yaml:"sms"`
}

type App struct {
	IsProduction bool          `yaml:"is-production"`
	RunMode      string        `yaml:"run-mode"`
	JwtSecret    string        `yaml:"jwt-secret"`
	Expiration   time.Duration `yaml:"expiration"`
	ApiVersion   string        `yaml:"api-version"`
	LogsSavePath string        `yaml:"logs-save-path"`
	OssUrlPrefix string        `yaml:"oss-url-prefix"`
	PwdSalt      string        `yaml:"pwd-salt"` //用户密码盐值
}

type Server struct {
	IsTls        bool          `yaml:"is-tls"`
	HttpPort     int           `yaml:"http-port"`
	ReadTimeOut  time.Duration `yaml:"read-time-out"`
	WriteTimeOut time.Duration `yaml:"write-time-out"`
	CertFile     string        `yaml:"cert-file"`
	KeyFile      string        `yaml:"key-file"`
}
