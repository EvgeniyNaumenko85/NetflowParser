package pkg

import (
	"NetflowParser/models"
	"NetflowParser/utilities"
	"encoding/binary"
	"fmt"
	"net"
)

func AcceptFlagsFromConsole() (FilePath string, NetFlowRecord models.NetFlowRecord, InputFields [4]int) {
	InputFields = [4]int{0, 0, 0, 0}

	FilePath = utilities.CheckFileExists()

	fmt.Println("нужен ли поиск по флагу `source:` 1 - да / нет?")
	var sourceField string
	fmt.Scan(&sourceField)
	if sourceField == "1" {
		NetFlowRecord.Source = net.ParseIP(utilities.InputCheckIP())
		InputFields[0] = 1
	}

	fmt.Println("нужен ли поиск по флагу `destination: ` 1 - да / нет?")
	var destinationField string
	fmt.Scan(&destinationField)
	if destinationField == "1" {
		NetFlowRecord.Destination = net.ParseIP(utilities.InputCheckIP())
		InputFields[1] = 1
	}

	fmt.Println("нужен ли поиск по флагу `account_id: ` 1 - да / нет?")
	var accountIDField string
	targetAccountID := make([]byte, 4, 4)
	fmt.Scan(&accountIDField)
	if accountIDField == "1" {
		targetAccountIDUint32 := utilities.InputCheckUint32()
		binary.LittleEndian.PutUint32(targetAccountID, targetAccountIDUint32)
		NetFlowRecord.AccountID = targetAccountID
		InputFields[2] = 1
	}

	fmt.Println("нужен ли поиск по флагу `tclass: ` 1 - да / нет?")
	var tclassField string
	targetTClass := make([]byte, 4, 4)
	fmt.Scan(&tclassField)
	if tclassField == "1" {
		targetTClassUint32 := utilities.InputCheckUint32()
		binary.LittleEndian.PutUint32(targetTClass, targetTClassUint32)
		NetFlowRecord.TClass = targetTClass
		InputFields[3] = 1
		fmt.Println("InputFields", InputFields)
	}

	return FilePath, NetFlowRecord, InputFields
}
