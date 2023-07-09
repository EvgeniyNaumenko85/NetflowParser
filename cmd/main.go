package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"runtime"
	"sync"
	"time"
)

type NetFlowRecord struct {
	Source      net.IP
	Destination net.IP
	AccountID   uint32
	TClass      uint32
}

func main() {

	startTime := time.Now()

	// Путь к бинарному файлу статистики NetFlow
	filePath := "C:\\Users\\Евгений Науменко\\Desktop\\задание Babylon\\iptraffic_raw_1688454355.utm"

	// Значения параметров командной строки для фильтрации
	//targetSource := net.ParseIP("127.0.0.1")
	//targetDestination := net.ParseIP("192.168.0.1")
	//targetAccountID := uint32(123)
	//targetTClass := uint32(1)

	//Значения параметров командной строки для фильтрации
	targetSource := net.ParseIP("")
	targetDestination := net.ParseIP("")
	targetAccountID := uint32(1)
	targetTClass := uint32(1)

	// Открытие файла в режиме чтения
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Пропуск двоичного заголовка файла (349 байт)
	headerSize := 349
	_, err = file.Seek(int64(headerSize), 0)
	if err != nil {
		fmt.Println("Ошибка при перемещении указателя файла:", err)
		return
	}

	// Количество горутин для параллельной обработки
	numWorkers := runtime.NumCPU()

	// Канал для передачи записей между горутинами и главной функцией
	recordChan := make(chan NetFlowRecord)

	// Ограничитель для дожидания завершения всех горутин
	wg := sync.WaitGroup{}
	wg.Add(numWorkers)

	// Запуск горутин для обработки записей
	for i := 0; i < numWorkers; i++ {
		go func() {
			for record := range recordChan {
				// Сравнение значений с заданными параметрами
				if record.Source.Equal(targetSource) && record.Destination.Equal(targetDestination) &&
					record.AccountID == targetAccountID && record.TClass == targetTClass {
					// Значения соответствуют фильтру - выполняем необходимые действия
					fmt.Printf("Найдена запись со значениями: source=%s, destination=%s, account_id=%d, tclass=%d\n",
						record.Source.String(), record.Destination.String(), record.AccountID, record.TClass)
				}
			}
			wg.Done()
		}()
	}

	// Чтение оставшихся данных из файла
	recordSize := 74
	recordCount := 0

	for {
		recordData := make([]byte, recordSize)
		_, err := file.Read(recordData)
		if err != nil {
			// Достигнут конец файла или произошла ошибка чтения
			break
		}

		record := parseNetFlowRecord(recordData)

		// Отправка записи в канал для обработки горутинами
		recordChan <- record

		recordCount++
	}

	// Закрытие канала после чтения всех записей
	close(recordChan)

	// Дожидаемся завершения работы всех горутин
	wg.Wait()

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	fmt.Printf("Прочитано %d записей за время %s\n", recordCount, elapsedTime)
}

func parseNetFlowRecord(data []byte) NetFlowRecord {
	record := NetFlowRecord{}

	record.Source = net.IP(data[1:5]) //todo проверить net.IP
	record.Destination = net.IP(data[5:9])
	record.AccountID = binary.LittleEndian.Uint32(data[53:57])
	record.TClass = binary.LittleEndian.Uint32(data[61:65])

	return record
}
