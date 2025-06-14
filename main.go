package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/fatih/color"
)

type IPStat struct {
    IP    string
    Count int
}

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
	ipCounts := make(map[string]int)
	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		ip,status,err := ParseLine(line)
		if err != nil{
			continue
		}
		ipCounts[ip] += 1
		statusCounts[status] += 1

		lineCount++
	}
	if err := scanner.Err(); err != nil{
		log.Fatal("Ошибка чтения файла:", err)
	}

	fmt.Printf("Всего строк: %d\n", lineCount)

	red := color.New(color.FgRed).PrintfFunc()
	green := color.New(color.FgGreen).PrintfFunc()

	fmt.Println("Статистика HTTP-статусов:")
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

	var ipStats []IPStat
	for ip, count := range ipCounts{
		ipStats = append(ipStats, IPStat{ip, count})
	}

	sort.Slice(ipStats, func(i, j int) bool {
		return ipStats[i].Count > ipStats[j].Count
	})

	fmt.Println("Top 5 IPs:")
	for i := 0; i < 5 && i < len(ipStats); i++ {
		red("%s: %d запросов\n", ipStats[i].IP, ipStats[i].Count)
	}
}