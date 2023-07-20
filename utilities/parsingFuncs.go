package utilities

import (
	"NetflowParser/models"
	"bytes"
	"sync"
	"sync/atomic"
)

func SelectParseFuncNumb(inputFields [4]int) (int, error) {
	var (
		digBase = 1
		counter int
	)

	for _, val := range inputFields {
		counter += digBase * val
		digBase *= 2
	}

	//todo включить проверку!
	//проверка наличия хотя бы одного флага
	if counter == 0 {
		return 0, models.ErrNoFlags
	}

	return counter, nil
}

// Определяем тип функции для проверки условия
type checkFunc func(models.NetFlowRecord, models.NetFlowRecord) bool

// Создаем слайс с функциями проверки
var checkFunctions = []checkFunc{
	func(record, netFlowRecord models.NetFlowRecord) bool {
		return record.Source.Equal(netFlowRecord.Source)
	},
	func(record, netFlowRecord models.NetFlowRecord) bool {
		return record.Destination.Equal(netFlowRecord.Destination)
	},
	func(record, netFlowRecord models.NetFlowRecord) bool {
		return record.Source.Equal(netFlowRecord.Source) && record.Destination.Equal(netFlowRecord.Destination)
	},
	func(record, netFlowRecord models.NetFlowRecord) bool {
		return bytes.Equal(record.AccountID, netFlowRecord.AccountID)
	},
	func(record, netFlowRecord models.NetFlowRecord) bool {
		return record.Source.Equal(netFlowRecord.Source) && bytes.Equal(record.AccountID, netFlowRecord.AccountID)
	},
	func(record, netFlowRecord models.NetFlowRecord) bool {
		return record.Destination.Equal(netFlowRecord.Destination) && bytes.Equal(record.AccountID, netFlowRecord.AccountID)
	},
	func(record, netFlowRecord models.NetFlowRecord) bool {
		return record.Source.Equal(netFlowRecord.Source) && record.Destination.Equal(netFlowRecord.Destination) && bytes.Equal(record.AccountID, netFlowRecord.AccountID)
	},
	func(record, netFlowRecord models.NetFlowRecord) bool {
		return bytes.Equal(record.TClass, netFlowRecord.TClass)
	},
	func(record, netFlowRecord models.NetFlowRecord) bool {
		return record.Source.Equal(netFlowRecord.Source) && bytes.Equal(record.TClass, netFlowRecord.TClass)
	},
	func(record, netFlowRecord models.NetFlowRecord) bool {
		return record.Destination.Equal(netFlowRecord.Destination) && bytes.Equal(record.TClass, netFlowRecord.TClass)
	},
	func(record, netFlowRecord models.NetFlowRecord) bool {
		return record.Source.Equal(netFlowRecord.Source) && record.Destination.Equal(netFlowRecord.Destination) && bytes.Equal(record.TClass, netFlowRecord.TClass)
	},
	func(record, netFlowRecord models.NetFlowRecord) bool {
		return bytes.Equal(record.AccountID, netFlowRecord.AccountID) && bytes.Equal(record.TClass, netFlowRecord.TClass)
	},
	func(record, netFlowRecord models.NetFlowRecord) bool {
		return record.Source.Equal(netFlowRecord.Source) && bytes.Equal(record.AccountID, netFlowRecord.AccountID) && bytes.Equal(record.TClass, netFlowRecord.TClass)
	},
	func(record, netFlowRecord models.NetFlowRecord) bool {
		return record.Destination.Equal(netFlowRecord.Destination) && bytes.Equal(record.AccountID, netFlowRecord.AccountID) && bytes.Equal(record.TClass, netFlowRecord.TClass)
	},
	func(record, netFlowRecord models.NetFlowRecord) bool {
		return record.Source.Equal(netFlowRecord.Source) && record.Destination.Equal(netFlowRecord.Destination) && bytes.Equal(record.AccountID, netFlowRecord.AccountID) && bytes.Equal(record.TClass, netFlowRecord.TClass)
	},
}

func Worker(recordChan chan models.NetFlowRecord, NetFlowRecord models.NetFlowRecord, wg *sync.WaitGroup, counterChan chan<- uint64, funcNumber int) {
	defer wg.Done()
	var counter uint64
	for record := range recordChan {
		// Вызываем нужную функцию проверки условия
		if checkFunctions[funcNumber-1](record, NetFlowRecord) {
			// todo добавляем данные в БД
			atomic.AddUint64(&counter, 1)
		}
	}
	counterChan <- counter
}
