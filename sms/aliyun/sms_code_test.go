package aliyun

import (
	"log"
	"testing"
)

func TestSendSMSCode(t *testing.T) {
	response, err := SendSMSCode("18842671207", "1266", Sllt)
	if err != nil {
		log.Println("err", err)
	}
	log.Println("response", response)
}
