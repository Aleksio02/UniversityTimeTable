package services

import (
	"auth/cmd/auth/connectors"
	stresponse "auth/cmd/auth/model/response"
	"auth/cmd/auth/utils"
)

func IsAuthorized(chatId int) bool {
	// TODO: alexeyi: implement me
	return false
}

func AuthenticateUser(tempCode string, chatId string) {
	githubTokenResponse, _ := connectors.GetUserTokenFromGithub(tempCode)
	jsonResponseBody := stresponse.GetGitHubUserToken{}
	utils.WriteBodyToObject(githubTokenResponse.Body, &jsonResponseBody)
	githubTokenResponse.Body.Close()

	// save user info into db if it not exists

	githubUserResponse, _ := connectors.GetUserInfo(jsonResponseBody.AccessToken, jsonResponseBody.TokenType)
	githubUserResponseBody := stresponse.GitHubGetUserResponse{}
	utils.WriteBodyToObject(githubUserResponse.Body, &githubUserResponseBody)
	githubUserResponse.Body.Close()

	var userInfo = stresponse.UserInfo{}
	userInfo.TelegramChatId = chatId
	userInfo.GitHubId = githubUserResponseBody.Id
	var authResponse = stresponse.AuthResponse{Status: 200, UserInfo: userInfo}

	connectors.SendUserInfo(authResponse)
}
