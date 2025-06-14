package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func ParseLine(line string) (ip string, status int, err error){
	parts := strings.Split(line, " ")

	if len(parts) < 9 {
		return "", 0, fmt.Errorf("неверный формат строки: ожидается не менее 9 частей")
	}

	ip_check := net.ParseIP(parts[0])
    if ip_check == nil {
        return "", 0, fmt.Errorf("невалидный IP-адрес: %s", parts[0])
    }

	status, err = strconv.Atoi(parts[8])
	if err != nil {
		return "", 0, fmt.Errorf("ошибка парсинга статуса: %v", err)
	}

	if status < 100 || status > 599 {
        return "", 0, fmt.Errorf("код статуса не входит в допустимую область: %d", status)
    }

	if strings.HasPrefix(parts[4], "\"") || strings.HasSuffix(parts[6], "\"") {
        return "", 0, fmt.Errorf("неправильно сформированная строка")
    }

	return parts[0], status, nil
}