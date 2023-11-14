package response

type AuthResponse struct {
	Status   int `json:"status"`
	Response any `json:"response"`
}

type GitHubGetUserResponse struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
}
