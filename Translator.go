package TranslatorAPI

import (
	"TranslatorAPI/services"
	"fmt"
	"log"
)

func Translator(platform, word, SorLang, TarLang string) {
	result, err := services.YoudaoSendRequests(word, SorLang, TarLang)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
