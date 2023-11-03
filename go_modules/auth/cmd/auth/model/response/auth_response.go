package response

type AuthResponse struct {
	Status          int    `json:"status"`
	ResponseMessage string `json:"responseMessage"`
	AuthLink        string `json:"authLink"`
}
