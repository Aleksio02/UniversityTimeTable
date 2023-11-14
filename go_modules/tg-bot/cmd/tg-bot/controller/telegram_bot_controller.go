package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tg-bot/cmd/tg-bot/bot"
	response "tg-bot/cmd/tg-bot/model/request"
	"tg-bot/cmd/tg-bot/utils"
)

func ReceiveAuthInfo(c *gin.Context) {
	requestBody := response.AuthInfoRequest{}
	utils.WriteBodyToObject(c.Request.Body, &requestBody)

	if requestBody.Status == 401 {
		//	send message about unauthorized with auth link
		authLink := requestBody.Response.(string)
		msg := tgbotapi.NewMessage(int64(requestBody.ChatId), fmt.Sprintf("Вы не зарегистрированы!\nДля прохождения регистрации пройдите по ссылке: %s", authLink))
		bot.Bot.Send(msg)

	} else if requestBody.Status == 200 {
		//  send message about successfully authorization
		msg := tgbotapi.NewMessage(int64(requestBody.ChatId), "Вы прошли авторизацию.")
		msg.ReplyMarkup = bot.BasicFeatures
		bot.Bot.Send(msg)
	}

}
