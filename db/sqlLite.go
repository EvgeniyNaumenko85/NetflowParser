package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// MySQLDBConnect Устанавливаем соединение с базой данных MySQL
func MySQLDBConnect() {
	db, err := sql.Open("mysql", "user:password@tcp(hostname:port)/database_name")
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return
	}

	// Проверка соединения с базой данных
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
}
