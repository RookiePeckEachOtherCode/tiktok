package model

// API的状态码和状态信息
type Response struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"`
}
