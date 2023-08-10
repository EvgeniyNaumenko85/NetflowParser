package pkg

import (
	"NetflowParser/common"
	"NetflowParser/models"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

func ResultToLog(fileName, filePath, logfilePath string, counter uint64, NFR models.NetFlowRecord, recordCount int, elapsedTime time.Duration) {
	// Функция для форматирования значений <nil> или 0
	printFormatter := func(value interface{}) string {
		if reflect.ValueOf(value).Kind() == reflect.Ptr {
			if reflect.ValueOf(value).IsNil() {
				return "<nil>"
			}
		} else if reflect.ValueOf(value).Kind() == reflect.Uint32 {
			if value.(uint32) == 0 {
				return "<nil>"
			}
		}
		return fmt.Sprintf("%v", value)
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Println("ошибка при получении информации о файле:", err)
		return
	}

	fileSizeBytes := fileInfo.Size()
	fileSizeMB := float64(fileSizeBytes) / (1024 * 1024)

	log.Printf("\n>---------------\nфайл: %s\nпо адресу: %s\nразмер файла: %.2f МБ\nНайдено записей: %d по флагам:\n\tsource "+
		"= %s\n\tdestination = %s\n\taccount_id = %s\n\ttclass = %s\nПрочитано: %d записей, за время: %s\n--------------->\n\n",
		fileName, filePath, fileSizeMB, counter,
		printFormatter(NFR.Source.String()), printFormatter(NFR.Destination.String()),
		printFormatter(common.BytesToUint32LE(NFR.AccountID)), printFormatter(common.BytesToUint32LE(NFR.TClass)),
		recordCount, elapsedTime)

	fmt.Printf("\n>---------------\nфайл: %s\nпо адресу: %s\nразмер файла: %.2f МБ\nНайдено записей: %d по флагам:\n\tsource"+
		" = %s\n\tdestination = %s\n\taccount_id = %s\n\ttclass = %s\nПрочитано: %d записей, за время: %s\n--------------->\n",
		fileName, filePath, fileSizeMB, counter,
		printFormatter(NFR.Source.String()), printFormatter(NFR.Destination.String()),
		printFormatter(common.BytesToUint32LE(NFR.AccountID)), printFormatter(common.BytesToUint32LE(NFR.TClass)),
		recordCount, elapsedTime)

	fmt.Println()
	fmt.Printf("Информация о результате работы можно посмотреть в файле: %s\n", logfilePath)
	fmt.Println("Press any key to quit...")
}
