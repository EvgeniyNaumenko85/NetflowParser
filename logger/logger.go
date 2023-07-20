package logger

import (
	"log"
	"os"
)

func LogginResult() (*os.File, error) {
	logfile, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Ошибка открытия файла лога: %v", err)
	}
	return logfile, nil
}
