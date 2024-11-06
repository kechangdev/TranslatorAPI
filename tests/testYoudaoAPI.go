package tests

import "TranslatorAPI"

func YoudaoAPI(word string) {
	TranslatorAPI.Translator("Youdao", word, "auto", "zh-CHS")
}
