package caiyun

// Request 发向彩云的请求
type Request struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserID    string `json:"user_id"`
}
