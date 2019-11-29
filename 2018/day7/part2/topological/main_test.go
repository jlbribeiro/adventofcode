package topological_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2018/day7/part2/topological"
)

func TestOrder(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		nWorkers int
		baseCost int
		expected int
	}{
		{
			name: "example",
			input: `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`,
			nWorkers: 2,
			baseCost: 0,
			expected: 15,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)

			actual, err := topological.Order(reader, tc.nWorkers, tc.baseCost)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
