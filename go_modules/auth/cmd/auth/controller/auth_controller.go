package controller

import (
	"auth/cmd/auth/config"
	"auth/cmd/auth/connectors"
	"auth/cmd/auth/model/request"
	stresponse "auth/cmd/auth/model/response"
	"auth/cmd/auth/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func GetSession(c *gin.Context) {
	requestBody := request.GetSessionRequest{}

	writeBodyToObject(c.Request.Body, &requestBody)

	var getSessionResponse stresponse.AuthResponse

	if isAuthorized(requestBody.ChatId) {
		// TODO: alexeyi: implement this part
		getSessionResponse.Status = 200
		getSessionResponse.ResponseMessage = "Trying authorize chatId - " + strconv.Itoa(requestBody.ChatId)
	} else {
		getSessionResponse.Status = 401
		getSessionResponse.AuthLink = generateAuthLink(requestBody.ChatId)
	}

	c.JSON(http.StatusOK, getSessionResponse)
}

func SendToken(c *gin.Context) {
	fullURL := fmt.Sprintf("%s://%s:%d%s",
		config.Config.Application.Protocol,
		config.Config.Application.Host,
		config.Config.Application.Port,
		c.Request.RequestURI)

	myUrl, _ := url.Parse(fullURL)
	pathParams, _ := url.ParseQuery(myUrl.RawQuery)
	res, _ := connectors.GetUserTokenFromGithub(pathParams.Get("code"))
	jsonResponseBody := stresponse.GetGitHubUserToken{}
	writeBodyToObject(res.Body, &jsonResponseBody)
	res.Body.Close()

	// save user info into db if it not exists
	chatId := pathParams["chatId"][0]

	githubUserResponse, _ := connectors.GetUserInfo(jsonResponseBody.AccessToken, jsonResponseBody.TokenType)
	githubUserResponseBody := stresponse.GitHubGetUserResponse{}
	writeBodyToObject(githubUserResponse.Body, &githubUserResponseBody)
	githubUserResponse.Body.Close()

	var userInfo = stresponse.UserInfo{}
	userInfo.TelegramChatId = chatId
	userInfo.GitHubId = githubUserResponseBody.Id
	var authResponse = stresponse.AuthResponse{Status: 200, UserInfo: userInfo}

	connectors.SendTokenInfo(authResponse)

	c.JSON(200, "You can close this page")
}

func MockTgBot(c *gin.Context) {
	jsonRequestBody := stresponse.AuthResponse{}
	writeBodyToObject(c.Request.Body, &jsonRequestBody)

	fmt.Println(jsonRequestBody)
}

func writeBodyToObject(body io.Reader, objectToWrite any) {
	responseBody, _ := ioutil.ReadAll(body)
	json.Unmarshal(responseBody, &objectToWrite)
}

func generateAuthLink(chatId int) string {
	githubHostIp := "https://github.com"
	githubAuthorizeMethod := "/login/oauth/authorize"
	clientId := config.Config.GitHub.ClientId
	redirectUri := fmt.Sprintf("%s://%s:%d/%s/sendToken",
		config.Config.Application.Protocol,
		config.Config.Application.Host,
		config.Config.Application.Port,
		config.Config.Application.Name)

	return fmt.Sprintf("%s%s?%s&%s?%s",
		githubHostIp,
		githubAuthorizeMethod,
		utils.CreatePathParam("client_id", clientId),
		utils.CreatePathParam("redirect_uri", redirectUri),
		utils.CreatePathParam("chatId", strconv.Itoa(chatId)))
}

func isAuthorized(chatId int) bool {
	// TODO: alexeyi: implement me
	return false
}
