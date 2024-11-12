package services

import (
	"TranslatorAPI"
	"fmt"
	"testing"
)

func TestYoudaoAPI(t *testing.T) {
	resq, err := TranslatorAPI.Done("Youdao", "test", "auto", "zh-CHS")
	if err != nil {
		return
	}
	fmt.Println(resq)
}
