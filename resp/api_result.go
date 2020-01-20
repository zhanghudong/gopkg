package resp

import (
	"fmt"
	"net/http"
)

/**
 * @Author: zhanghudong
 * @Date: 2019-11-03 20:19
 */

type ApiResult struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Details    IApiErr     `json:"details,omitempty"`
	RequestId  string      `json:"request_id,omitempty"`
	Data       interface{} `json:"data"`
}

func (a ApiResult) Error() string {
	return fmt.Sprintf("StatusCode=%d, Message=%s, Details=%v", a.StatusCode, a.Message, a.Details)
}

func Ok(data interface{}) *ApiResult {
	return &ApiResult{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		Data:       data,
	}
}

func NewResult(code int, message string) *ApiResult {
	return &ApiResult{
		StatusCode: code,
		Message:    message,
	}
}

func BadRequestDetails(apiErr *ApiError) *ApiResult {
	return &ApiResult{
		StatusCode: http.StatusBadRequest,
		Message:    http.StatusText(http.StatusBadRequest),
		Details:    apiErr,
	}
}

func BadRequestApiErr(code int64, msg string) *ApiResult {
	return &ApiResult{
		StatusCode: http.StatusBadRequest,
		Message:    http.StatusText(http.StatusBadRequest),
		Details: &ApiError{
			ErrCode: code,
			ErrMsg:  msg,
		},
	}
}

func ForbiddenDetails(details string) *ApiResult {
	return &ApiResult{
		StatusCode: http.StatusForbidden,
		Message:    details,
	}
}

var (
	//从客户端发送的请求在服务端被正常处理了。
	ApiOK = NewResult(http.StatusOK, http.StatusText(http.StatusOK))

	//服务器接受的请求已成功处理，但返回的响应报文的主体部分不包含实体（浏览器页面不更新，仅发送信息给服务器）
	ApiErrNoContent = NewResult(http.StatusNoContent, http.StatusText(http.StatusNoContent))

	//请求报文中存在语法错误
	ApiErrBadRequest = NewResult(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))

	//请求需要有通过HTTP认证的认证信息。另外如果之前已进行一次请求，则表示用户认证失败
	ApiErrUnauthorized = NewResult(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))

	//对请求资源的访问被服务器拒绝了，如未获得文件系统的访问授权，访问权限出现某些问题
	ApiErrForbidden = NewResult(http.StatusForbidden, http.StatusText(http.StatusForbidden))

	//服务器上没有请求的资源
	ApiErrNotFound = NewResult(http.StatusNotFound, http.StatusText(http.StatusNotFound))

	//服务器在执行请求时发生了错误，Bug或临时故障
	ApiErrInternalServerError = NewResult(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))

	//服务器暂时处于超负荷或正在进行停机维护，现在无法处理请求
	ApiErrServiceUnavailable = NewResult(http.StatusServiceUnavailable, http.StatusText(http.StatusServiceUnavailable))
)
