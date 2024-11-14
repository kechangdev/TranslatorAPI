package translator

import (
	"TranslatorAPI/internal/models/translator"
	"errors"
)

// 定义通用错误
var (
	ErrUnsupportedPlatform = errors.New("unsupported translation platform")
	ErrEmptyText           = errors.New("empty text")
	ErrInvalidLanguage     = errors.New("invalid language code")
)

// Translator 翻译器接口
type Translator interface {
	// Translate 执行翻译
	// text: 要翻译的文本
	// from: 源语言代码
	// to: 目标语言代码
	// 返回翻译结果和可能的错误
	Translate(text, from, to string) (translator.TranslateResponse, error)

	// GetName 获取翻译器名称
	GetName() string

	// IsSupported 检查是否支持指定的语言
	IsSupported(lang string) bool
}
