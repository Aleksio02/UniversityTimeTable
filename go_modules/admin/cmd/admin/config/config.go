package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Config appConfig

type appConfig struct {
	// Пример переменной, загружаемой в функции LoadConfig
	Application struct {
		Name    string
		Version string
		Port    int
	}
}

// LoadConfig загружает конфигурацию из файлов
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("application") // <- имя конфигурационного файла
	v.SetConfigType("yml")
	v.SetEnvPrefix("blueprint")
	v.AutomaticEnv()
	for _, path := range configPaths {
		v.AddConfigPath(path) // <- // путь для поиска конфигурационного файла в
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	return v.Unmarshal(&Config)
}
