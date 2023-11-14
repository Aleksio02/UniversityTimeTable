package connectors

import (
	"auth/cmd/auth/model/response"
	"bytes"
	"encoding/json"
	"net/http"
)

func SendUserInfo(authResponse response.AuthResponse) (*http.Response, error) {
	methodName := "/sendAuthInfo"
	requestURL := TELEGRAM_BOT_HOST_IP + methodName
	requestURL = ""

	authResponseAsReader, _ := json.Marshal(authResponse)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", requestURL, bytes.NewBuffer(authResponseAsReader))

	req.Header.Set("Accept", "application/json")
	return client.Do(req)
}

// TODO: alexeyi: set port to tg-bot and change it here
const TELEGRAM_BOT_HOST_IP string = "http://localhost:8080/tg-bot"
