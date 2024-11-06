package TranslatorAPI

import (
	"TranslatorAPI/services"
	"fmt"
	"log"
)

func Translator(platform, word, SorLang, TarLang string) {
	switch platform {
	case "Youdao":
		result, err := services.YoudaoSendRequests(word, SorLang, TarLang)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	case "Caiyun":
		result, err := services.CaiyunSendRequests(word, SorLang, TarLang)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(word, "UK:", result.Dictionary.Prons.En, "US:", result.Dictionary.Prons.EnUs)
		for _, item := range result.Dictionary.Explanations {
			fmt.Println(item)
		}
	}

}
