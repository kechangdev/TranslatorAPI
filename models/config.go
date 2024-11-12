package models

type Config struct {
	Youdao YoudaoConfig `json:"Youdao"`
	Other  OtherConfig  `json:"other"`
}

type YoudaoConfig struct {
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
}

type OtherConfig struct {
	// 添加其他字段
}
