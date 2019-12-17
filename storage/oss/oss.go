package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zhanghudong/gopkg/logger"
)

type OssConfig struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyId     string `yaml:"access-key-id"`
	AccessKeySecret string `yaml:"access-key-secret"`
	BucketName      string `yaml:"bucket-name"`
}

type alyOss struct {
	Bucket *oss.Bucket
}

var _alyOss *alyOss

func InitOssClient(config *OssConfig) error {
	// 创建OSSClient实例
	client, err := oss.New(config.Endpoint, config.AccessKeyId, config.AccessKeySecret)
	if err != nil {
		logger.Sugar.Errorf("initOssClient oss.New err=%s", err.Error())
		return err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(config.BucketName)
	if err != nil {
		logger.Sugar.Errorf("initOssClient client.Bucket err=%s", err.Error())
		return err
	}
	_alyOss = &alyOss{
		Bucket: bucket,
	}
	return nil
}
