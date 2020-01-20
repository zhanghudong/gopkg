package captcha

import (
	"github.com/mojocn/base64Captcha"
	"github.com/zhanghudong/gopkg/logger"
)

//configJsonBody json request body.
type Config struct {
	CaptchaType   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

type Result struct {
	Id   string
	Data string
}

var store = base64Captcha.DefaultMemStore

// base64Captcha create http handler
func GenerateCaptcha(param Config) (*Result, error) {
	var driver base64Captcha.Driver

	switch param.CaptchaType {
	case "audio":
		driver = param.DriverAudio
	case "string":
		driver = param.DriverString.ConvertFonts()
	case "math":
		driver = param.DriverMath.ConvertFonts()
	case "chinese":
		driver = param.DriverChinese.ConvertFonts()
	default:
		driver = param.DriverDigit
	}
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	if err != nil {
		return nil, err
	}
	return &Result{
		Id:   id,
		Data: b64s,
	}, nil
}

func Verify(id, data string) bool {
	logger.Sugar.Infof("id=%s, data=%s, getData=%s", id, data, store.Get(id, false))
	if store.Verify(id, data, true) {
		return true
	}
	return false
}
