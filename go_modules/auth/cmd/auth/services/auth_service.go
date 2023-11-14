package services

import (
	"auth/cmd/auth/connectors"
	"auth/cmd/auth/dao"
	"auth/cmd/auth/model"
	stresponse "auth/cmd/auth/model/response"
	"auth/cmd/auth/utils"
	"errors"
	"strconv"
)

func CreateUser(tempCode string, chatId string) {
	/// Getting user token from github
	githubTokenResponse, _ := connectors.GetUserTokenFromGithub(tempCode)
	githubTokenResponseBody := stresponse.GetGitHubUserToken{}
	utils.WriteBodyToObject(githubTokenResponse.Body, &githubTokenResponseBody)
	githubTokenResponse.Body.Close()

	/// Getting information about GitHub account by token
	githubUserResponse, _ := connectors.GetUserInfo(githubTokenResponseBody.AccessToken, githubTokenResponseBody.TokenType)
	githubUserResponseBody := stresponse.GitHubGetUserResponse{}
	utils.WriteBodyToObject(githubUserResponse.Body, &githubUserResponseBody)
	githubUserResponse.Body.Close()

	/// Saving user info with given chat id and GitHub account id
	var userInfo = model.User{}
	userInfo.TelegramChatId, _ = strconv.Atoi(chatId)
	userInfo.GithubUserId = githubUserResponseBody.Id
	userInfo = dao.CreateUser(userInfo)

	var authResponse = stresponse.AuthResponse{Status: 200, Response: userInfo}

	connectors.SendUserInfo(authResponse)
}

func GetUserByTelegramChatId(chatId int) (model.User, error) {
	foundUser := dao.GetUserByTelegramChatId(chatId)
	if foundUser.TelegramChatId == 0 {
		return model.User{}, errors.New("User not found!")
	}
	return foundUser, nil
}
