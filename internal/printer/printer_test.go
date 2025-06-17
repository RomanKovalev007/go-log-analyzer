package printer

import (
	"testing"

	"github.com/RomanKovalev007/go-log-analyzer/internal/aggregator"
)



func TestPrinter(t *testing.T) {
    stats := aggregator.Stats{
        StatusCounts: map[int]int{200: 1, 404: 1},
        IPCounts:     map[string]int{"127.0.0.1": 1},
    }
    p := NewPrinter(5)
    p.PrintResult(stats) // Проверяем, что не падает
}