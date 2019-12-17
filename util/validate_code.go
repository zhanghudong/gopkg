package util

import (
	"fmt"
	"github.com/zhanghudong/gopkg/logger"
	"math/rand"
	"strings"
	"time"
)

func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, err := fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
		if err != nil {
			logger.Sugar.Errorf("GenValidateCode err=%s", err)
		}
	}
	str := sb.String()
	if str[0:1] == "0" {
		return GenValidateCode(width)
	}
	return sb.String()
}
