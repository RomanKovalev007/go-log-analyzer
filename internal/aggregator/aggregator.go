package aggregator

import parcer "github.com/RomanKovalev007/go-log-analyzer/internal/parser"

type Stats struct{
	IPCounts map[string]int
	StatusCounts map[int]int
}

type Aggregator struct{
	stats Stats
}

func NewAggregator() *Aggregator {
	return &Aggregator{stats: Stats{
		StatusCounts: make(map[int]int),
		IPCounts: make(map[string]int),
	}}
}

func (a *Aggregator) Add(entry parcer.LogInfo) {
	a.stats.StatusCounts[entry.Status] ++
	a.stats.IPCounts[entry.IP] ++
}

func (a *Aggregator) Results() Stats{
	return a.stats
}