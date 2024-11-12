package services

import (
	"TranslatorAPI"
	"fmt"
	"testing"
)

func TestCaiyunAPI(t *testing.T) {
	resq, err := TranslatorAPI.Done("Caiyun", "test", "ZN", "US")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resq)
}
