package connectors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tg-bot/cmd/tg-bot/utils"
)

func AuthenticateUser(chatId int) {
	methodName := "/getSession"
	//requestURL := AUTH_HOST_IP + methodName
	requestURL := fmt.Sprintf("%s%s?%s", AUTH_HOST_IP, methodName, utils.CreatePathParam("chatId", strconv.Itoa(chatId)))
	//requestURL = ""

	type RequestBody struct {
		chatId int `json:"chatId"`
	}

	requestBodyAsReader, _ := json.Marshal(RequestBody{chatId: chatId})

	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURL, bytes.NewBuffer(requestBodyAsReader))

	req.Header.Set("Accept", "application/json")
	client.Do(req)
}

const AUTH_HOST_IP string = "http://localhost:8081/auth"
