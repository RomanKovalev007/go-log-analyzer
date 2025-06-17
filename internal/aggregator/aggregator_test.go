package aggregator

import (
	"testing"

	"github.com/RomanKovalev007/go-log-analyzer/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestAggregator(t *testing.T){
	tests := []struct{
		name string
		agg Aggregator
		addStats []parser.LogInfo
		resultStat Stats
	}{
		{
			name: "first test",
			agg: *NewAggregator(),
			addStats: []parser.LogInfo{
				{IP: "127.0.0.1", Status: 404},
				{IP: "192.168.1.1", Status: 404},
				{IP: "127.0.0.1", Status: 403},
				{IP: "2001:db8::1", Status: 301},
				{IP: "127.0.0.1", Status: 200},
				{IP: "192.168.1.1", Status: 200},
				{IP: "127.0.0.1", Status: 200},
				{IP: "192.168.1.1", Status: 503},
			},
			resultStat: Stats{
				IPCounts: map[string]int{
					"127.0.0.1": 4,
					"192.168.1.1": 3,
					"2001:db8::1": 1,
				},
				StatusCounts: map[int]int{
					200: 3,
					301: 1,
					403: 1,
					404: 2,
					503: 1,
				},
			},
		},
				{
			name: "second test",
			agg: *NewAggregator(),
			addStats: []parser.LogInfo{
				{IP: "2001:db8::1", Status: 200},
				{IP: "192.168.1.1", Status: 404},
				{IP: "127.0.0.1", Status: 403},
				{IP: "2001:db8::1", Status: 302},
				{IP: "127.0.0.1", Status: 200},
				{IP: "2001:db8::1", Status: 403},
				{IP: "127.0.0.1", Status: 200},
				{IP: "192.168.1.1", Status: 503},
			},
			resultStat: Stats{
				IPCounts: map[string]int{
					"127.0.0.1": 3,
					"192.168.1.1": 2,
					"2001:db8::1": 3,
				},
				StatusCounts: map[int]int{
					200: 3,
					302: 1,
					403: 2,
					404: 1,
					503: 1,
				},
			},
		},
	}

	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T) {
			test_agg := NewAggregator()
			for _, stat := range tt.addStats{
				test_agg.Add(stat)
			}
			assert.Equal(t, test_agg.stats.IPCounts, tt.resultStat.IPCounts, "Для переданных IP адресов неверно посчитано их количество")
			assert.Equal(t, test_agg.stats.StatusCounts, tt.resultStat.StatusCounts, "Для переданных статусов неверно посчитано их количество")
		})
	}
}