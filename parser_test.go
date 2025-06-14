package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParceLine_Valid(t *testing.T) {
	tests := []struct{
		name string
		input string
		expectedIP string
		expectedStatus int
	}{
		{
			name: "Nginx Standart Log",
			input: `127.0.0.1 - - [10/Oct/2023:13:55:36 +0000] "GET /api HTTP/1.1" 200 1234`,
			expectedIP: "127.0.0.1",
			expectedStatus: 200,
		},
		{
			name: "Apache Combined Log",
			input: `192.168.1.1 - user [10/Oct/2023:14:01:23 +0000] "POST /login HTTP/1.0" 404 345`,
			expectedIP: "192.168.1.1",
			expectedStatus: 404,
		},
		{
			name: "IPv6 Address",
			input: `2001:db8::1 - - [10/Oct/2023:15:00:01 +0000] "HEAD /status HTTP/1.1" 503 0`,
			expectedIP: "2001:db8::1",
			expectedStatus: 503,
		},
		{
			name: "With Query Parameters",
			input: `127.0.0.1 - - [10/Oct/2023:16:10:45 +0000] "GET /search?q=go HTTP/1.1" 200 5120`,
			expectedIP: "127.0.0.1",
			expectedStatus: 200,
		},
	}

	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T) {
			ip, status, err := ParseLine(tt.input)
			require.NoError(t, err, "Ошибка парсинга валидной строки")
			assert.Equal(t, tt.expectedIP, ip, "Неверный IP-адрес")
			assert.Equal(t, tt.expectedStatus, status, "Неверный статус")
		})
	}
}

func TestParceLine_InValid(t *testing.T) {
	tests := []struct{
		name string
		input string
	}{
		{"Empty String", ""},
		{"Only IP", "127.0.0.1"},
		{"Invalid IP", "300.0.0.1 - - [...]"},
		{"Missing Status", `127.0.0.1 - - [10/Oct] "GET /"`},
		{"Invalid Status", `127.0.0.1 - - [10/Oct] "GET /" ABC 123`},
		{"Malformed Request", `127.0.0.1 - - [10/Oct] "GET / HTTP/1.1"`},
	}
	for _,tt := range tests{
		t.Run(tt.name, func(t *testing.T) {
			_,_,err := ParseLine(tt.input)
			assert.Error(t, err, "ожидалась ошибка для невалидного формата")
		})
	}
}
