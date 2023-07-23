package main

import (
	"NetflowParser/db"
	"NetflowParser/logger"
	"NetflowParser/models"
	"NetflowParser/pkg"
	"NetflowParser/utilities"
	"fmt"
	"github.com/spf13/viper"
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
		log.Fatalf("ошибка открытия файла лога: %v", err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	err = pkg.LoadEnvVariables()
	if err != nil {
		log.Fatalf("ошибка открытия файла лога: %v", err)
		return
	}

	//Секрет
	isValid, err := utilities.CheckSecretExpiration()
	if err != nil {
		log.Println("ошибка при проверке даты:", err)
		return
	}
	if !isValid {
		log.Println("секрет не валиден")
		return
	}

	config := db.Config{
		DriverName: viper.GetString("dbMySQL.driver"),
		User:       viper.GetString("dbMySQL.user"),
		Password:   os.Getenv("DB_PASSWORD"),
		Protocol:   viper.GetString("dbMySQL.protocol"),
		Host:       viper.GetString("dbMySQL.host"),
		Port:       viper.GetString("dbMySQL.port"),
		DBName:     viper.GetString("dbMySQL.dbname"),
	}

	//Подключаем БД
	database := db.StartDbConnection(config)

	fmt.Print("поднимаем таблицу в БД: ")
	if err = db.Up(database); err != nil {
		fmt.Printf("Error while migrating tables, err is: %s", err.Error())
		log.Fatalf("Error while migrating tables, err is: %s", err.Error())
		return
	}
	fmt.Print("успешно")
	fmt.Println()

	// Канал для сигнализации об окончании обработки
	done := make(chan string)

	//Graceful shutdown
	go func() {

		// Канал для сигнализации об окончании обработки
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		<-done
		<-quit

		db.CloseDbConnection(config)

		log.Println("Shutting down")
		os.Exit(0)
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
		fmt.Println("ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Пропуск двоичного заголовка файла (349 байт)
	headerSize := 349
	err = pkg.OmittingFileBinHeader(headerSize, file)
	if err != nil {
		fmt.Println("ошибка при перемещении указателя файла:", err)
		return
	}

	// Количество горутин для параллельной обработки
	numWorkers := runtime.NumCPU() * 35 //todo проверить оптимальное кол-во. горутин

	// Канал для передачи записей между горутинами и главной функцией
	recordChan := make(chan models.NetFlowRecord, numWorkers*1000)

	// Канал для передачи количества найденныхзаписей между горутинами и главной функцией
	counterChan := make(chan uint64, numWorkers*1000)

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

	pkg.ResultToLog(fileName, filePath, logfilePath, counter, NetFlowRecord, recordCount, elapsedTime)

	exit := ""
	fmt.Scan(&exit)
	done <- exit
}
