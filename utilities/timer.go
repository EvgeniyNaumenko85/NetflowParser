package utilities

import (
	"time"
)

var endDate = "2023-09-26"

// CheckSecretExpiration Функция для проверки срока действия секрета
func CheckSecretExpiration() (bool, error) {
	secretDate, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return false, err
	}

	currentTime := time.Now()
	if currentTime.After(secretDate) {
		return false, nil // Дата истекла
	}

	return true, nil
}
