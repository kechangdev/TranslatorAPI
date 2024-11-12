package tests

import (
	"TranslatorAPI"
	"fmt"
	"testing"
)

var word = "test"

func TestCaiyunAPI(t *testing.T) {
	resq, err := TranslatorAPI.Translator("Caiyun", word, "ZN", "US")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resq)
}
