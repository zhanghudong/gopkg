package aliyun

import (
	"log"
	"testing"
)

func TestSendSMSCode(t *testing.T) {
	response, err := SendSMSCode("18842671208", "1266", SLLT)
	if err != nil {
		log.Println("err", err)
	}
	log.Println("response", response)
}
