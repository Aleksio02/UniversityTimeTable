package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tg-bot/cmd/tg-bot/connectors"
)

func StartBotHandler(tgBot *tgbotapi.BotAPI, updateConfig tgbotapi.UpdateConfig) {
	updates := tgBot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		msg.Text = update.Message.Text
		if update.Message.Text == "/start" {
			msg.Text = "Добро пожаловать! Начинаем процесс аутентификации"
			sendMessage(tgBot, msg)
			handleStart(update.Message.Chat.ID)
		} else if update.Message.Text == "/numeric" {
			msg.ReplyMarkup = NumericKeyboard
			sendMessage(tgBot, msg)
		} else if update.Message.Text == "/romanian" {
			msg.ReplyMarkup = RomanianKeyboard
			sendMessage(tgBot, msg)
		} else if update.Message.Text == "/close" {
			tgbotapi.NewRemoveKeyboard(true)
			msg.ReplyMarkup = nil
		}

	}
}

func handleStart(id int64) {
	connectors.AuthenticateUser(int(id))
}

func sendMessage(tgBot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	if _, err := tgBot.Send(msg); err != nil {
		panic(err)
	}
}
