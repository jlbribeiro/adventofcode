package freq_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2018/day1/part2/freq"
)

func TestAnalyse(t *testing.T) {
	var testCases = []struct {
		name     string
		deltas   []int
		expected int
	}{
		{
			name:     "description",
			deltas:   []int{+1, -2, +3, +1, +1, -2},
			expected: 2,
		},
		{
			name:     "example1",
			deltas:   []int{+1, -1},
			expected: 0,
		},
		{
			name:     "example2",
			deltas:   []int{3, +3, +4, -2, -4},
			expected: 10,
		},
		{
			name:     "example3",
			deltas:   []int{-6, +3, +8, +5, -6},
			expected: 5,
		},
		{
			name:     "example4",
			deltas:   []int{+7, +7, -2, -7, -4},
			expected: 14,
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
