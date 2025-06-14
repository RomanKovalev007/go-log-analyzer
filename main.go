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
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		lineCount++
	}
	if err := scanner.Err(); err != nil{
		log.Fatal("Ошибка чтения файла:", err)
	}
	fmt.Printf("Всего строк: %d\n", lineCount)
}