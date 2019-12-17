package resp

/**
 * @Author: zhanghudong
 * @Date: 2019-11-03 20:17
 */

type ApiError struct {
	ErrCode int64  `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}
