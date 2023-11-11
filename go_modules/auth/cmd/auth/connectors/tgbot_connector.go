package connectors

import (
	"auth/cmd/auth/model/response"
	"bytes"
	"encoding/json"
	"net/http"
)

// TODO: alexeyi: set port to tg-bot and change it here
const TELEGRAM_BOT_HOST_IP string = "http://localhost:8080/tg-bot"

func SendUserInfo(authResponse response.AuthResponse) (*http.Response, error) {
	methodName := "/sendAuthInfo"
	requestURL := TELEGRAM_BOT_HOST_IP + methodName
	requestURL = ""

	authResponseAsReader, _ := json.Marshal(authResponse)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", requestURL, bytes.NewReader(authResponseAsReader))
	req.Header.Set("Accept", "application/json")
	// todo: alexeyi: send user info as request body
	return client.Do(req)
}
