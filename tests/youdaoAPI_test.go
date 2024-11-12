package tests

import "TranslatorAPI"

func YoudaoAPI(word string) (TranslatorAPI.TResponse, error) {
	return TranslatorAPI.Translator("Youdao", word, "auto", "zh-CHS")
}
