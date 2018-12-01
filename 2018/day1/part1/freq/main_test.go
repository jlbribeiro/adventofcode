package freq_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2018/day1/part1/freq"
)

func TestAnalyse(t *testing.T) {
	var testCases = []struct {
		name     string
		deltas   []int
		expected int
	}{
		{
			name:     "simple",
			deltas:   []int{+1, -2, +3, +1},
			expected: 3,
		},
		{
			name:     "example1",
			deltas:   []int{+1, +1, +1},
			expected: 3,
		},
		{
			name:     "example2",
			deltas:   []int{+1, +1, -2},
			expected: 0,
		},
		{
			name:     "example3",
			deltas:   []int{-1, -2, -3},
			expected: -6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := freq.Analyse(tc.deltas)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
