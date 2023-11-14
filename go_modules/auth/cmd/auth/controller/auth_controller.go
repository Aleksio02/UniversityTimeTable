package controller

import (
	"auth/cmd/auth/config"
	"auth/cmd/auth/model/request"
	stresponse "auth/cmd/auth/model/response"
	"auth/cmd/auth/services"
	"auth/cmd/auth/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
)

func GetSession(c *gin.Context) {
	requestBody := request.GetSessionRequest{}

	utils.WriteBodyToObject(c.Request.Body, &requestBody)
	var getSessionResponse stresponse.AuthResponse
	foundUser, err := services.GetUserByTelegramChatId(requestBody.ChatId)

	if err == nil {
		getSessionResponse.Status = 200
		getSessionResponse.Response = foundUser
	} else {
		getSessionResponse.Status = 401
		getSessionResponse.Response = generateAuthLink(requestBody.ChatId)
	}

	c.JSON(http.StatusOK, getSessionResponse)
}

func CreateUser(c *gin.Context) {
	fullURL := fmt.Sprintf("%s://%s:%d%s",
		config.Config.Application.Protocol,
		config.Config.Application.Host,
		config.Config.Application.Port,
		c.Request.RequestURI)
	myUrl, _ := url.Parse(fullURL)
	pathParams, _ := url.ParseQuery(myUrl.RawQuery)

	services.CreateUser(pathParams.Get("code"), pathParams["chatId"][0])
	c.JSON(200, "You can close this page")
}

// TODO: alexeyi: remove it
func MockTgBot(c *gin.Context) {
	jsonRequestBody := stresponse.AuthResponse{}
	utils.WriteBodyToObject(c.Request.Body, &jsonRequestBody)
	fmt.Println("")
}

func generateAuthLink(chatId int) string {
	githubHostIp := "https://github.com"
	githubAuthorizeMethod := "/login/oauth/authorize"
	clientId := config.Config.GitHub.ClientId
	redirectUri := fmt.Sprintf("%s://%s:%d/%s/createUser",
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
