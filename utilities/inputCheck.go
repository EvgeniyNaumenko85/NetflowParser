package utilities

import (
	"fmt"
	"strconv"
	"strings"
)

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
