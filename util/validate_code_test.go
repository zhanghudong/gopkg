package util

import (
	"log"
	"testing"
)

func TestGenValidateCode(t *testing.T) {
	for i := 0; i < 100000; i++ {
		log.Println(GenValidateCode(6))
	}

}
