package tests

import "TranslatorAPI"

func CaiyunAPI(word string) {
	TranslatorAPI.Translator("Caiyun", word, "ZN", "US")
}
