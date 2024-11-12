package tests

import (
	"TranslatorAPI"
	"fmt"
	"testing"
)

func TestYoudaoAPI(t *testing.T) {
	resq, err := TranslatorAPI.Translator("Youdao", "test", "auto", "zh-CHS")
	if err != nil {
		return
	}
	fmt.Println(resq)
}
