package parser

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

var (
	ErrInvalidFormat = errors.New("invalid log format")
	ErrInvalidIP     = errors.New("invalid IP address")
	ErrInvalidStatus = errors.New("invalid HTTP status")
)

type LogInfo struct{
	IP string
	Status int
}

func ParseLine(line string) (LogInfo, error){
	parts := strings.Split(line, " ")

	if len(parts) < 9 {
		return LogInfo{}, fmt.Errorf("%v: ожидается не менее 9 частей", ErrInvalidFormat)
	}

	ip_check := net.ParseIP(parts[0])
    if ip_check == nil {
        return LogInfo{}, fmt.Errorf("%v: %s",ErrInvalidIP, parts[0])
    }

	status, err := strconv.Atoi(parts[8])
	if err != nil {
		return LogInfo{}, fmt.Errorf("%v: %v",ErrInvalidStatus, err)
	}

	if status < 100 || status > 599 {
        return LogInfo{}, fmt.Errorf("%v: код статуса не входит в допустимую область: %d",ErrInvalidStatus, status)
    }

	return LogInfo{parts[0], status}, nil
}