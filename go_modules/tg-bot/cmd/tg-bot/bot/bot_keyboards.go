package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var BasicFeatures = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Просмотреть расписание"),
		tgbotapi.NewKeyboardButton("Добавить информацию о себе"),
		tgbotapi.NewKeyboardButton("Изменить расписание"),
	),
)

var NumericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
	),
)

var RomanianKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("I"),
		tgbotapi.NewKeyboardButton("II"),
		tgbotapi.NewKeyboardButton("III"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("IV"),
		tgbotapi.NewKeyboardButton("V"),
		tgbotapi.NewKeyboardButton("VI"),
	),
)
