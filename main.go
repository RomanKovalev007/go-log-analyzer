package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("logs/access.log")
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

	for status, counts := range statusCounts{
		fmt.Printf("статус %d: %d запросов\n", status, counts)
	}
}