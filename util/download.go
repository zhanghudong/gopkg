package util

import (
	"fmt"
	"github.com/zhanghudong/gopkg/logger"
	"io/ioutil"
	"net/http"
)

//图片格式
const (
	ImageJpegContentType = "image/jpeg"
	imagePngContentType  = "image/png"
)

func GetDownloadFile(url string, fileType string) ([]byte, error) {
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		logger.Sugar.Errorf("GetDownloadFile http.DefaultClient.Get err=%s", err.Error())
		return nil, err
	}
	if fileType != "" {
		if fileType != resp.Header.Get("Content-Type") {
			return nil, fmt.Errorf("response Content-Type not %s,Content-Type=%s", fileType, resp.Header.Get("Content-Type"))
		}
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Sugar.Errorf("GetDownloadFile ioutil.ReadAll err=%s", err.Error())
		return nil, err
	}
	if resp.StatusCode != 200 {
		logger.Sugar.Errorf("GetDownloadFile resp.StatusCode =%d", resp.StatusCode)
		return nil, fmt.Errorf("response StatusCode=%d,Message=%s", resp.StatusCode, string(bytes))
	}
	return bytes, err
}

func GetDownloadImage(url string) ([]byte, string, int64, error) {
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		logger.Sugar.Errorf("GetDownloadFile http.DefaultClient.Get err=%s", err.Error())
		return nil, "", 0, err
	}
	cType := resp.Header.Get("Content-Type")
	if ImageJpegContentType != cType && imagePngContentType != cType {
		return nil, "", 0, fmt.Errorf("response Content-Type not image type,Content-Type=%s", resp.Header.Get("Content-Type"))
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Sugar.Errorf("GetDownloadFile ioutil.ReadAll err=%s", err.Error())
		return nil, "", 0, err
	}
	if resp.StatusCode != 200 {
		logger.Sugar.Errorf("GetDownloadFile resp.StatusCode =%d", resp.StatusCode)
		return nil, "", 0, fmt.Errorf("response StatusCode=%d,Message=%s", resp.StatusCode, string(bytes))
	}
	return bytes, cType, resp.ContentLength, nil
}
