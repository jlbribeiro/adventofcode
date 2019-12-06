package thermal_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day5/part1/thermal"
)

func TestRun(t *testing.T) {
	var testCases = []struct {
		name     string
		program  []int
		expected []int
	}{
		{
			name:     "example1",
			program:  []int{},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := thermal.Run(tc.program, 1)
			if len(actual) != len(tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
			for i := range tc.expected {
				if actual[i] != tc.expected[i] {
					t.Errorf("expected %v, got %v", tc.expected, actual)
					return
				}
			}
		})
	}
}
