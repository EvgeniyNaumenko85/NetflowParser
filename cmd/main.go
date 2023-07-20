package main

import (
	"NetflowParser/logger"
	"NetflowParser/models"
	"NetflowParser/pkg"
	"NetflowParser/utilities"
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func main() {

	//Логирование
	logfile, err := logger.LogginResult()
	if err != nil {
		log.Fatalf("Ошибка открытия файла лога: %v", err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	//Секрет
	isValid, err := utilities.CheckSecretExpiration()
	if err != nil {
		log.Println("Ошибка при проверке секретной даты:", err)
		return
	}
	if !isValid {
		log.Println("Секрет не валиден")
		return
	}

	// Канал для сигнализации об окончании обработки
	done := make(chan string)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	//graceful shutdown
	go func() {
		<-done
		<-quit

		//if err := db.CloseDbConnection(); err != nil {
		//	fmt.Errorf("error occurred on database connection closing: %s", err.Error())
		//}

		os.Exit(0) // Завершить программу с кодом 0 (успешное завершение)
		//log.Println("Shutting down")
		//os.Exit(0) // Завершить программу с кодом 0 (успешное завершение)
	}()

	// Принимаем флаги из консоли
	filePath, NetFlowRecord, InputFields := pkg.AcceptFlagsFromConsole()

	// Выбираем функцию для парсинга
	funcNumber, err := utilities.SelectParseFuncNumb(InputFields)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Открытие файла в режиме чтения
	file, err := pkg.OpenFileByFilePath(filePath)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Пропуск двоичного заголовка файла (349 байт)
	headerSize := 349
	err = pkg.OmittingFileBinHeader(headerSize, file)
	if err != nil {
		fmt.Println("Ошибка при перемещении указателя файла:", err)
		return
	}

	// Количество горутин для параллельной обработки
	numWorkers := runtime.NumCPU() * 100 //todo проверить оптимальное кол-во. горутин

	// Канал для передачи записей между горутинами и главной функцией
	recordChan := make(chan models.NetFlowRecord, numWorkers*10)

	// Канал для передачи количества найденныхзаписей между горутинами и главной функцией
	counterChan := make(chan uint64, numWorkers*10)

	// Ограничитель для дожидания завершения всех горутин
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Запуск учета времени обработки файла
	startTime := time.Now()

	// Запуск горутины-работники
	for i := 0; i < numWorkers; i++ {
		go utilities.Worker(recordChan, NetFlowRecord, &wg, counterChan, funcNumber)
	}

	// Чтение оставшихся данных из файла
	recordSize := 74
	recordCount := 0

	for {
		recordData := make([]byte, recordSize)
		_, err = file.Read(recordData)
		if err != nil {
			// Достигнут конец файла или произошла ошибка чтения
			break
		}

		record := pkg.ParseNetFlowRecord(recordData)

		// Отправка записи в канал для обработки горутинами
		recordChan <- record
		recordCount++
	}

	// Закрытие канала после чтения всех записей
	close(recordChan)

	// Дожидаемся завершения работы всех горутин
	wg.Wait()
	close(counterChan)

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	// Создаем переменную для хранения общего количества найденных записей
	var counter uint64

	// Получаем общее количество найденных
	for val := range counterChan {
		counter += val
	}

	filePath = file.Name()
	fileName := filepath.Base(file.Name())
	logfilePath := logfile.Name()

	log.Printf(">---------------\nфайл: %s\nпо адресу: %s\nНайдено записей: %d со значениями:\n\tsource = "+
		"%s\n\tdestination = %s\n\taccount_id = %d\n\ttclass = %d\nПрочитано: %d записей, за время: %s\n--------------->\n",
		fileName, filePath, counter, NetFlowRecord.Source.String(), NetFlowRecord.Destination.String(),
		utilities.BytesToUint32LE(NetFlowRecord.AccountID), utilities.BytesToUint32LE(NetFlowRecord.TClass),
		recordCount, elapsedTime)

	fmt.Printf("Информация о результате рботы можно посмотреть в файле: %s\n", logfilePath)
	fmt.Println("Pres any key to quit...")

	exit := ""
	fmt.Scan(&exit)
	done <- exit
}
