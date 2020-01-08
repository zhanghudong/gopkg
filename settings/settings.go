package settings

import (
	"github.com/zhanghudong/gopkg/db"
	"github.com/zhanghudong/gopkg/logger"
	"github.com/zhanghudong/gopkg/sms/aliyun"
	"github.com/zhanghudong/gopkg/storage/oss"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

/**
 * @Author: zhanghudong
 * @Date: 2019-12-13 18:05
 */

func loadConfig(filename string, config interface{}) error {
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(in, config)
}

func ConfInit(options ...string) {
	var err error
	//生成配置默认值

	for i, v := range options {
		if i == 0 {
			ApplicationConfig = &Config{}
			if err = loadConfig(v, ApplicationConfig); err != nil {
				panic(err)
			}
			if ApplicationConfig.App.LogsSavePath == "" {
				panic("logs save path does not set")
			}
			if err = logger.InitLogger(ApplicationConfig.App.LogsSavePath); err != nil {
				panic(err)
			}
			continue
		}
		switch v {
		case ALiYunOss:
			if ApplicationConfig.Oss == nil {
				panic("oss config  does not set")
			}
			err = oss.InitOssClient(ApplicationConfig.Oss)
			if err != nil {
				panic(err)
			}
		case ALiYunSms:
			if ApplicationConfig.Sms == nil {
				panic("sms config  does not set")
			}
			err = aliyun.InitAliYunSms(ApplicationConfig.Sms)
			if err != nil {
				panic(err)
			}
		case Mysql:
			if ApplicationConfig.Mysql == nil {
				panic("sms config  does not set")
			}
			err = db.InitMysql(ApplicationConfig.Mysql)
			if err != nil {
				panic(err)
			}
		default:
			panic("options does not exist")
		}
	}
}
