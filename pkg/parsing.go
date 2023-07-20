package pkg

import (
	"NetflowParser/models"
	"os"
)

func OpenFileByFilePath(filePath string) (file *os.File, err error) {
	file, err = os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func OmittingFileBinHeader(headerSize int, file *os.File) error {
	_, err := file.Seek(int64(headerSize), 0)
	if err != nil {
		return err
	}
	return nil
}

func ParseNetFlowRecord(data []byte) models.NetFlowRecord {
	record := models.NetFlowRecord{}

	record.Source = data[1:5] //
	//fmt.Println("record.Source: ", record.Source)
	record.Destination = data[5:9]
	//fmt.Println("record.Destination: ", record.Destination)
	record.AccountID = data[53:57]
	//fmt.Println("record.AccountID: ", record.AccountID)
	record.TClass = data[61:65]
	//fmt.Println("record.TClass: ", record.TClass)
	//fmt.Println()

	return record
}
