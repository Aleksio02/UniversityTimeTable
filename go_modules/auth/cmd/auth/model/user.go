package model

type User struct {
	TelegramChatId int      `json:"telegramchatid"`
	GithubUserId   int      `json:"githubuserid"`
	Roles          []string `json:"roles"`
	AdditionalData struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
	} `json:"additionaldata"`
}
