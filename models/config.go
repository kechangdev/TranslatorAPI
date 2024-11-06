package models

type Config struct {
	Youdao YoudaoConfig `json:"Youdao"`
	Other  OtherConfig  `json:"other"` // 如果有其他配置项
}

type YoudaoConfig struct {
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
}

// 如果有其他配置项，可以定义 OtherConfig 结构体
type OtherConfig struct {
	// 添加其他字段
}
