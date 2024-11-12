package tests

import "TranslatorAPI"

func CaiyunAPI(word string) (TranslatorAPI.TResponse, error) {
	return TranslatorAPI.Translator("Caiyun", word, "ZN", "US")
}
