package resp

import "encoding/json"

/**
 * @Author: zhanghudong
 * @Date: 2019-11-03 20:17
 */

type ApiError struct {
	ErrCode int64  `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

func NewApiErr(code int64, msg string) *ApiError {
	return &ApiError{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (a ApiError) ApiErr() string {
	bytes, err := json.Marshal(a)
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

type IApiErr interface {
	ApiErr() string
}
