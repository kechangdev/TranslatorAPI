package translator

// Response  翻译得到的响应
type Response struct {
	Word      []string  `json:"word"`
	SoundMark SoundMark `json:"soundMark"`
}
