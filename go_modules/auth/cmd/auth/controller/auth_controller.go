package controller

import (
	"auth/cmd/auth/config"
	"auth/cmd/auth/connectors"
	stresponse "auth/cmd/auth/model/response"
	"auth/cmd/auth/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetSession(c *gin.Context) {
	//chatId, _ := strconv.ParseInt(c.Param("chatId"), 10, 32)
	requestBody := getResponseBody(c.Request.Body)

	var getSessionResponse stresponse.AuthResponse

	if isAuthorized(requestBody["chatId"]) {
		// TODO: alexeyi: implement this part
		getSessionResponse.Status = 200
		getSessionResponse.ResponseMessage = "Trying authorize chatId - " + requestBody["chatId"]
	} else {
		getSessionResponse.Status = 401
		getSessionResponse.AuthLink = generateAuthLink(requestBody["chatId"])
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
	jsonResponseBody := getResponseBody(res.Body)
	res.Body.Close()

	// save user info into db if it not exists
	chatId := pathParams["chatId"]
	fmt.Println(chatId)

	githubUserResponse, _ := connectors.GetUserInfo(jsonResponseBody["access_token"], jsonResponseBody["token_type"])
	githubUserResponseBody := getGitHubUserInfo(githubUserResponse.Body)

	var userInfo = stresponse.UserInfo{}
	userInfo.GitHubId = githubUserResponseBody.Id
	var authResponse = stresponse.AuthResponse{Status: 200, UserInfo: userInfo}

	connectors.SendTokenInfo(authResponse)

	c.JSON(200, "You can close this page")
}

func MockTgBot(c *gin.Context) {
	jsonRequestBody := getAuthResponse(c.Request.Body)

	fmt.Println(jsonRequestBody)
}

// TODO: alexeyi: it should be 1 function
func getResponseBody(body io.Reader) map[string]string {
	responseBody, _ := ioutil.ReadAll(body)
	jsonResponseBody := map[string]string{}
	json.Unmarshal(responseBody, &jsonResponseBody)
	return jsonResponseBody
}

func getGitHubUserInfo(body io.Reader) stresponse.GitHubGetUserResponse {
	responseBody, _ := ioutil.ReadAll(body)
	jsonResponseBody := stresponse.GitHubGetUserResponse{}
	json.Unmarshal(responseBody, &jsonResponseBody)
	return jsonResponseBody
}

func getAuthResponse(body io.Reader) stresponse.AuthResponse {
	responseBody, _ := ioutil.ReadAll(body)
	jsonResponseBody := stresponse.AuthResponse{}
	json.Unmarshal(responseBody, &jsonResponseBody)
	return jsonResponseBody
}

//////////////////////////////////////

func generateAuthLink(chatId string) string {
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
		utils.CreatePathParam("chatId", chatId))
}

func isAuthorized(chatId string) bool {
	// TODO: alexeyi: implement me
	return false
}
