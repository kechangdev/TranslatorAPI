package TranslatorAPI

import (
	"TranslatorAPI/services"
	"errors"
	"fmt"
	"log"
)

type SoundMark struct {
	UK string `json:"uk"`
	US string `json:"us"`
}
type TResponse struct {
	Word      []string  `json:"word"` // Translated Word
	SoundMark SoundMark `json:"soundMark"`
}

func Translator(platform, word, SorLang, TarLang string) (TResponse, error) {
	var resp TResponse
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
		resp.Word = result.Dictionary.Explanations
		resp.SoundMark.UK = result.Dictionary.Prons.En
		resp.SoundMark.US = result.Dictionary.Prons.EnUs
		return resp, nil
	}
	return resp, errors.New("Void Response")
}
