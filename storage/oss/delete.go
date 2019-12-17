package oss

import (
	"github.com/zhanghudong/gopkg/logger"
)

func Delete(key string) error {
	// 创建OSSClient实例。
	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err := _alyOss.Bucket.DeleteObject(key)
	if err != nil {
		logger.Sugar.Errorf("bucket.DeleteObject err=%s", err.Error())
		return err
	}
	return nil
}
