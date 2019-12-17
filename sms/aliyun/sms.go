package aliyun

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/zhanghudong/gopkg/logger"
)

type SmsConfig struct {
	AccessKey    string `yaml:"access-key"`
	AccessSecret string `yaml:"access-secret"`
	SignName     string `yaml:"sign-name"`
}

var _alySms *SmsConfig

func InitAliYunSms(config *SmsConfig) error {
	_alySms = &SmsConfig{
		AccessKey:    config.AccessKey,
		AccessSecret: config.AccessSecret,
		SignName:     config.SignName,
	}
	return nil
}

const (
	SLLT = "SMS_180340999"
)

type SMSResponse struct {
	Message   string `json:"Message"`
	RequestId string `json:"RequestId"`
	BizId     string `json:"BizId"`
	Code      string `json:"Code"`
}

type SMSCode struct {
	Phone     string `json:"phone"`
	Code      string `json:"code"`
	ExpiresIn int64  `json:"expires_in"`
}

func SendSMSCode(phone, code, templateCode string) (*SMSResponse, error) {
	client, err := sdk.NewClientWithAccessKey("default", _alySms.AccessKey, _alySms.AccessSecret)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["PhoneNumbers"] = phone
	request.QueryParams["TemplateCode"] = templateCode
	request.QueryParams["SignName"] = _alySms.SignName
	request.QueryParams["TemplateParam"] = "{\"code\":" + code + "}"
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		logger.Sugar.Error("----", err)
		logger.Sugar.Error("----", response)
		return nil, err
	}
	sMSResponse := SMSResponse{}
	err = json.Unmarshal([]byte(response.GetHttpContentString()), &sMSResponse)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	return &sMSResponse, nil
}

func SendSMSINTCode(phone, code, templateCode string) (*SMSResponse, error) {
	client, err := sdk.NewClientWithAccessKey("default", _alySms.AccessKey, _alySms.AccessSecret)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["PhoneNumbers"] = phone
	request.QueryParams["TemplateCode"] = templateCode
	request.QueryParams["SignName"] = _alySms.SignName
	request.QueryParams["TemplateParam"] = "{\"code\":" + code + "}"
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		logger.Sugar.Error("----", err)
		logger.Sugar.Error("----", response)
		return nil, err
	}
	sMSResponse := SMSResponse{}
	err = json.Unmarshal([]byte(response.GetHttpContentString()), &sMSResponse)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	return &sMSResponse, nil
}
