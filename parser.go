package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseLine(line string) (ip string, status int, err error){
	parts := strings.Split(line, " ")

	if len(parts) < 10 {
		return "", 0, fmt.Errorf("неверный формат строки")
	}

	ip = parts[0]
	status, err = strconv.Atoi(parts[8])
	if err != nil {
		return "", 0, fmt.Errorf("ошибка парсинга статуса: %v", err)
	}

	return ip, status, nil
}