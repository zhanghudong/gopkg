package util

import (
	"log"
	"testing"
)

func TestGetDownloadFile(t *testing.T) {
	bytes, err := GetDownloadFile("https://wx.qlogo.cn/mmopen/vi_32/0TZVeBT7Y4wQyXWFFriaYyyLdRStfDdp4Gxjg2gJxptG5Qlj6gfYPdBDqsX4fXSShiakxNmxdjv6MwYqaicjGyfUw/132", "image/jpeg")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(bytes)
}
