package main

import (
	"fmt"
	"tg-bot/cmd/tg-bot/config"
)

func main() {
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}
	go config.StartHttpServer()

	// Starting tg-bot
	config.StartTelegramBot()
}
