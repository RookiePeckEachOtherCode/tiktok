package dao

type Response struct {
	StatusCode int    `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"`
}
