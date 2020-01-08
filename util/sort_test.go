package util

import (
	"log"
	"testing"
)

func TestStrToSort(t *testing.T) {

	strs := []string{
		"!#@#$",
		"abc",
		"ABC",
		"啊啊",
		"阿偶",
		"我靠",
		"你的",
		"呀哈",
		"234",
		"73哈",
		"爱你123",
		"哈哈的",
		"adj  ant",
		"203183",
		"…/n\n…*   ……*……*",
		",.?;d",
		"!#@#$",
		"……*……*……*",
	}

	for _, str := range strs {
		log.Println(str, " = ", StrToSort(str))
	}
}
