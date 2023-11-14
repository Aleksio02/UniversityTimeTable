package response

type AuthInfoRequest struct {
	Status   int `json:"status"`
	ChatId   int `json:"chatId"`
	Response any `json:"response"`
}
