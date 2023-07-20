package utilities

import (
	"bufio"
	"os"
	"strings"
	"time"
)

// ReadSecret Функция для чтения значений из файла секретов
func CheckSecretExpiration() (bool, error) {
	file, err := os.Open(".env")
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) == 2 && parts[0] == "endDate" {
			secretDateStr := parts[1]
			secretDate, err := time.Parse("2006-01-02", secretDateStr)
			if err != nil {
				return false, err
			}
			currentTime := time.Now()
			if currentTime.After(secretDate) {
				return false, nil // Дата истекла
			}
		} else {
			return false, err
		}
	}
	return true, nil
}
