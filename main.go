package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"

	"github.com/RomanKovalev007/go-log-analyzer/internal/aggregator"
	"github.com/RomanKovalev007/go-log-analyzer/internal/parser"
	"github.com/RomanKovalev007/go-log-analyzer/internal/printer"
	"github.com/RomanKovalev007/go-log-analyzer/internal/reader"
)

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory Usage: Alloc = %.2f MB", float64(m.Alloc)/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func LogFile(
	reader *reader.Buffer,
	aggregator aggregator.Aggregator,
	filePath *string) error{
		err := reader.ReadFile(filePath, func(s string) error {
		loginfo, err := parser.ParseLine(s)
		if err != nil{
			return err
		}
		aggregator.Add(loginfo)
		return nil
	})
	return err
}

func main() {
	filePath := flag.String("f", "logs/access.log", "Путь к лог-файлу")
	flag.Parse() 

	reader := reader.NewReader(4*1024)
	aggregator := aggregator.NewAggregator()
	printer := printer.NewPrinter(5)
	
	
	printMemUsage()
	if err := LogFile(reader, *aggregator, filePath); err != nil{
		log.Fatalf("File reading error: %v", err)
	}

	printer.PrintResult(aggregator.Results())
	printMemUsage()
}