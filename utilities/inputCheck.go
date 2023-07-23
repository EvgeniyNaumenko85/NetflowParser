package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CheckFileExists() string {
	var filePath string

	for {
		fmt.Print("введите адрес файла или директории: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			filePath = strings.TrimSpace(scanner.Text())

			_, err := os.Stat(filePath)
			if !os.IsNotExist(err) {
				return filePath
			}
			fmt.Printf("файл или директория '%s' не существует или адрес некорректен. Пожалуйста, проверьте введенный адрес.\n", filePath)
		} else {
			fmt.Println("ошибка при считывании ввода. Пожалуйста, повторите попытку.")
		}
	}
}

func InputCheckIP() string {
	var inputSource string

	for {
		fmt.Print("введите IP-адрес в формате IPv4:  ")
		fmt.Scan(&inputSource)

		if isValidIPv4(inputSource) {
			break
		}

		fmt.Println("некорректный IP-адрес. Пожалуйста, введите IP-адрес в формате IPv4 (111.111.111.111): ")
	}

	return inputSource
}

func isValidIPv4(ip string) bool {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return false
	}

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}

	return true
}

func InputCheckUint32() uint32 {
	var userInput string
	for {
		fmt.Println("введите флаг: ")
		fmt.Scan(&userInput)

		_, err := strconv.ParseUint(userInput, 10, 32)
		if err != nil {
			fmt.Println("некорректный флаг, пожалуйста, введите флаг в формате числа")
		} else {
			break
		}
	}

	targetAccountIDUint64, _ := strconv.ParseUint(userInput, 10, 32)
	return uint32(targetAccountIDUint64)
}
