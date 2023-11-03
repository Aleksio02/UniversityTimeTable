package connectors

import (
	"auth/cmd/auth/config"
	"auth/cmd/auth/utils"
	"fmt"
	"net/http"
)

const GITHUB_HOST_IP string = "https://github.com"

func GetUserTokenFromGithub(code string) (*http.Response, error) {
	getTokenMethod := "/login/oauth/access_token"

	clientId := config.Config.GitHub.ClientId
	clientSecret := config.Config.GitHub.Secret

	redirectUri := fmt.Sprintf("%s://%s:%d/%s",
		config.Config.Application.Protocol,
		config.Config.Application.Host,
		config.Config.Application.Port,
		config.Config.Application.Name)

	requestURL := fmt.Sprintf("%s%s?%s&%s&%s&%s",
		GITHUB_HOST_IP,
		getTokenMethod,
		utils.CreatePathParam("client_id", clientId),
		utils.CreatePathParam("client_secret", clientSecret),
		utils.CreatePathParam("code", code),
		utils.CreatePathParam("redirect_uri", redirectUri))

	client := &http.Client{}
	req, _ := http.NewRequest("POST", requestURL, nil)
	req.Header.Set("Accept", "application/json")
	return client.Do(req)
}
