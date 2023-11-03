package connectors

import (
	"net/http"
)

// TODO: alexeyi: set port to tg-bot and change it here
const TELEGRAM_BOT_HOST_IP string = "http://localhost:8080/tg-bot"

func SendTokenInfo(tokenInfo map[string]string) (*http.Response, error) {
	methodName := "/sendAuthInfo"
	requestURL := TELEGRAM_BOT_HOST_IP + methodName
	requestURL = ""

	client := &http.Client{}
	req, _ := http.NewRequest("POST", requestURL, nil)
	req.Header.Set("Accept", "application/json")
	// todo: alexeyi: send token into request body
	return client.Do(req)
}
