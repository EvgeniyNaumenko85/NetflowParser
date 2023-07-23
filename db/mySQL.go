package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var database *sql.DB

type Config struct {
	DriverName string
	User       string
	Password   string
	DBName     string
	Protocol   string
	Host       string
	Port       string
	Param      string
}

// MySQLDBConnect Устанавливаем соединение с базой данных MySQL
func MySQLDBConnect(cfg Config) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		cfg.User, cfg.Password, cfg.Protocol, cfg.Host, cfg.Port, cfg.DBName)
	db, err := sql.Open(cfg.DriverName, dsn)
	if err != nil {
		fmt.Println("ошибка при подключении к базе данных: ", err)
		log.Fatal("ошибка при подключении к базе данных: ", err.Error())
	}
	// Проверка соединения с базой данных
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("подключение к БД: успешно")
	return db
}

func StartDbConnection(cfg Config) *sql.DB {
	database = MySQLDBConnect(cfg)
	return database
}

func CloseDbConnection(cfg Config) {
	database = MySQLDBConnect(cfg)
	err := database.Close()
	if err != nil {
		fmt.Println("ошибка при закрытии соединения с базой данных:", err)
	}
	fmt.Println("соединение с БД закрыто.")
}

func GetDBConn() *sql.DB {
	return database
}
