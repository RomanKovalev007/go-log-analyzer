package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
)

func main() {
	filePath := flag.String("f", "logs/access.log", "Путь к лог-файлу")
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatal("Ошибка открытия файла:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	statusCounts := make(map[int]int)
	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		_,status,err := ParseLine(line)
		if err != nil{
			continue
		}
		statusCounts[status] += 1

		lineCount++
	}
	if err := scanner.Err(); err != nil{
		log.Fatal("Ошибка чтения файла:", err)
	}

	fmt.Printf("Всего строк: %d\n", lineCount)

	red := color.New(color.FgRed).PrintfFunc()
	green := color.New(color.FgGreen).PrintfFunc()

	fmt.Println("Статистика HTTP-статусов")
	for status, count := range statusCounts{
		if status >= 500 {
			red("%d: %d (ошибка сервера)\n", status, count)
		} else if status >= 400 {
			red("%d: %d (ошибка клиента)\n", status, count)
		} else if status >= 300{
			green("%d: %d (перенаправление))\n", status, count)
		} else{
			green("%d: %d (успех)\n", status, count)
		}
	}
}