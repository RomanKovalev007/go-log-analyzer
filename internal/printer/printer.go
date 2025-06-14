package printer

import (
	"fmt"
	"sort"
	"github.com/RomanKovalev007/go-log-analyzer/internal/aggregator"
	"github.com/fatih/color"
)

type Printer struct{
	topN int
}

func NewPrinter(topN int) *Printer{
	return &Printer{topN: topN}
}

func (p *Printer) PrintResult(stats aggregator.Stats){
	p.StatusResult(stats.StatusCounts)
	p.IPResults(stats.IPCounts)
}

func (p *Printer) StatusResult(counts map[int]int){
	red := color.New(color.FgRed).PrintfFunc()
	green := color.New(color.FgGreen).PrintfFunc()
	yellow := color.New(color.FgYellow).PrintfFunc()
	fmt.Println("HTTP Status Statistics:")
	for status, count := range counts {
		switch {
		case status >= 500:
			red("%d: %d (server error)\n", status, count)
		case status >= 400:
			red("%d: %d (client error)\n", status, count)
		case status >= 300:
			yellow("%d: %d (redirect)\n", status, count)
		default:
			green("%d: %d (success)\n", status, count)
		}
	}
}

func (p *Printer) IPResults(counts map[string]int){
	type ipStat struct {
		ip    string
		count int
	}

	ips := make([]ipStat, 0, len(counts))
	for ip, count := range counts {
		ips = append(ips, ipStat{ip, count})
	}

	sort.Slice(ips, func(i, j int) bool {
		return ips[i].count > ips[j].count
	})

	fmt.Printf("\nTop %d IPs:\n", p.topN)
	red := color.New(color.FgRed).PrintfFunc()
	for i := 0; i < p.topN && i < len(ips); i++ {
		red("%s: %d requests\n", ips[i].ip, ips[i].count)
	}
}