package response

type UserInfo struct {
	GitHubId       int
	TelegramUserId string
	TelegramChatId string
}

type AuthResponse struct {
	Status          int      `json:"status"`
	ResponseMessage string   `json:"responseMessage"`
	AuthLink        string   `json:"authLink"`
	UserInfo        UserInfo `json:"UserInfo"`
}

type GitHubGetUserResponse struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
}
