package common

import (
	"NetflowParser/db"
	"NetflowParser/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func AddRecordToDB(record models.NetFlowRecord) error {
	tx, err := db.GetDBConn().Begin()
	if err != nil {
		return fmt.Errorf("ошибка при начале транзакции: %v", err)
	}

	query := "INSERT INTO record (source, destination, account_id, tclass) VALUES (?, ?, ?, ?)"
	_, err = tx.Exec(query, record.Source.String(), record.Destination.String(),
		BytesToUint32LE(record.AccountID), BytesToUint32LE(record.TClass))
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("ошибка при выполнении SQL-запроса: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("ошибка при завершении транзакции: %v", err)
	}

	return nil
}

func AddRecordToDBWithRetry(record models.NetFlowRecord, maxRetries int) error {
	for i := 0; i < maxRetries; i++ {
		err := AddRecordToDB(record)
		if err == nil {
			return nil // Успешная запись, выходим из цикла
		}

		// Печатаем ошибку и повторяем попытку записи
		fmt.Printf("Ошибка записи в БД, повторная попытка (%d/%d): %v\n", i+1, maxRetries, err)
		time.Sleep(time.Second) // Пауза перед следующей попыткой (предотвращает бесконечные попытки при критической ошибке)
	}

	return fmt.Errorf("превышено количество попыток записи в БД")
}
