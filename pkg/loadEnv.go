package pkg

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

func LoadEnvVariables() error {

	if err := initConfig(); err != nil {
		fmt.Printf("ошибка инициализации конфигов: %s\n", err)
		log.Printf("ошибка инициализации конфигов: %s\n", err)
		log.Fatalf("ошибка инициализации конфигов: %s", err)
		return err
	}

	if err := godotenv.Load(); err != nil {
		fmt.Printf("ошибка загрузки окружающих переменных: %s", err)
		log.Printf("ошибка загрузки окружающих переменных: %s", err)
		log.Fatalf("ошибка загрузки окружающих переменных: %s", err)
		return err
	}
	return nil
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
