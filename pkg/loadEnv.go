package pkg

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

func LoadEnvVariables() error {

	if err := initConfig(); err != nil {
		log.Fatalf("ошибка инициализации конфигов: %s", err.Error())
		return err
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("ошибка загрузки окружающих переменных: %s", err.Error())
		return err
	}

	return nil
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
