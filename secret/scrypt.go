package secret

import (
	"fmt"
	"github.com/zhanghudong/gopkg/settings"
	"golang.org/x/crypto/scrypt"
)

//存储密码
func ScryptPwd(pwd string) (string, error) {
	dk, err := scrypt.Key([]byte(pwd), []byte(settings.ApplicationConfig.App.PwdSalt), 16384, 8, 1, 32)
	return fmt.Sprintf("%x", dk), err
}
