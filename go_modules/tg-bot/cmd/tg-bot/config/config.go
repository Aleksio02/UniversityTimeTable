package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
	"tg-bot/cmd/tg-bot/bot"
	"tg-bot/cmd/tg-bot/controller"
)

var Config appConfig

var Bot *tgbotapi.BotAPI
var err error

type appConfig struct {
	// Пример переменной, загружаемой в функции LoadConfig
	Application struct {
		Name     string
		Version  string
		Port     int
		Host     string
		Protocol string
	}
	Telegram struct {
		Token string
	}
}

// LoadConfig загружает конфигурацию из файлов
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("application") // <- имя конфигурационного файла
	v.SetConfigType("yml")
	v.SetEnvPrefix("utt")
	v.AutomaticEnv()
	for _, path := range configPaths {
		v.AddConfigPath(path) // <- // путь для поиска конфигурационного файла в
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	return v.Unmarshal(&Config)
}

func StartHttpServer() {
	r := gin.New()

	v1 := r.Group(fmt.Sprintf("/%s", Config.Application.Name))
	{
		//v1.Use(auth())
		v1.GET("/system/test", controller.SystemTest)
	}
	r.Run(fmt.Sprintf(":%v", Config.Application.Port))
}

func StartTelegramBot() {
	Bot, err = tgbotapi.NewBotAPI(Config.Telegram.Token)
	if err != nil {
		panic(err)
	}

	Bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	bot.StartBotHandler(Bot, updateConfig)
}
