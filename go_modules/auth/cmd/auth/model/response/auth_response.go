package response

type AuthResponse struct {
	Status   int `json:"status"`
	ChatId   int `json:"chatId"`
	Response any `json:"response"`
}

type GitHubGetUserResponse struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
}
