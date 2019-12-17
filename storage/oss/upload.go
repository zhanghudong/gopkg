package oss

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zhanghudong/gopkg/logger"
	"io"
	"strings"
)

func UploadString(key, fileString string) error {
	// 指定存储类型为归档存储。
	storageType := oss.ObjectStorageClass(oss.StorageArchive)
	// 指定访问权限为公共读。
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)
	// 上传字符串。
	err := _alyOss.Bucket.PutObject(key, strings.NewReader(fileString), storageType, objectAcl)
	if err != nil {
		logger.Sugar.Errorf("UploadString bucket.PutObject err=%s", err.Error())
		return err
	}
	return nil
}

func UploadBytes(key string, file []byte) error {
	// 上传Byte数组。
	err := _alyOss.Bucket.PutObject(key, bytes.NewReader(file))
	if err != nil {
		logger.Sugar.Errorf("UploadBytes bucket.PutObject err=%s", err.Error())
		return err
	}
	return nil
}

func UploadStream(key string, reader io.Reader) error {
	// 上传文件流。
	err := _alyOss.Bucket.PutObject(key, reader)
	if err != nil {
		logger.Sugar.Errorf("UploadLocal bucket.PutObject err=%s", err.Error())
		return err
	}
	return nil

}

func UploadLocal(key string, filePath string) error {
	// 上传本地文件。
	err := _alyOss.Bucket.PutObjectFromFile(key, filePath)
	if err != nil {
		logger.Sugar.Errorf("UploadLocal bucket.PutObject err=%s", err.Error())
		return err
	}
	return nil
}
