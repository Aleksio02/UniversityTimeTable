package config

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Config appConfig

var DbManager *mongo.Database

type appConfig struct {
	// Пример переменной, загружаемой в функции LoadConfig
	Application struct {
		Name     string
		Version  string
		Port     int
		Host     string
		Protocol string
	}
	GitHub struct {
		ClientId string
		Secret   string
	}
	Database struct {
		Host     string
		Port     int
		Username string
		Password string
		DbName   string
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

func CreateDatabaseConnection() {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", Config.Database.Host, Config.Database.Port))
	clientOptions.SetAuth(options.Credential{Username: Config.Database.Username, Password: Config.Database.Password})
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	DbManager = client.Database(Config.Database.DbName)
}
