package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sort"
	"strings"

	"github.com/fatih/color"
)

type IPStat struct {
    IP    string
    Count int
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory Usage:\n")
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func main() {
	filePath := flag.String("f", "logs/access.log", "Путь к лог-файлу")
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatal("Ошибка открытия файла:", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := make([]byte, 4*1024)

	statusCounts := make(map[int]int)
	ipCounts := make(map[string]int)
	lineCount := 0

	printMemUsage()

	var leftover string
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		chunk := leftover + string(buffer[:n])
		lines := strings.Split(chunk, "\n")
		leftover = lines[len(lines)-1]

		for _, line := range lines[:len(lines)-1]{
			ip,status,err := ParseLine(line)
			if err != nil{
				continue
		}
		ipCounts[ip] += 1
		statusCounts[status] += 1

		lineCount++
		}
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

	printMemUsage()
}