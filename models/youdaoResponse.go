package models

type YoudaoResponse struct {
	TSpeakURL     string   `json:"tSpeakUrl"`
	RequestID     string   `json:"requestId"`
	Query         string   `json:"query"`
	Translation   []string `json:"translation"`
	MTerminalDict struct {
		URL string `json:"url"`
	} `json:"mTerminalDict"`
	ErrorCode string `json:"errorCode"`
	Dict      struct {
		URL string `json:"url"`
	} `json:"dict"`
	Webdict struct {
		URL string `json:"url"`
	} `json:"webdict"`
	L        string `json:"l"`
	IsWord   bool   `json:"isWord"`
	SpeakURL string `json:"speakUrl"`
}
