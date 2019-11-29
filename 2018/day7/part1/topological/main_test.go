package topological_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2018/day7/part1/topological"
)

func TestOrder(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		expected string
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
			expected: "CABDFE",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := topological.Order(strings.NewReader(tc.input))
			if actual != tc.expected {
				t.Errorf("expected '%s', got '%s'", tc.expected, actual)
			}
		})
	}
}
