package gravity_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day2/part2/gravity"
)

func TestRun(t *testing.T) {
	var testCases = []struct {
		name     string
		program  []int
		expected int
	}{
		{
			name:     "example1",
			program:  []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			expected: 3500,
		},
		{
			name:     "example2",
			program:  []int{1, 0, 0, 0, 99},
			expected: 2,
		},
		{
			name:     "example3",
			program:  []int{2, 3, 0, 3, 99},
			expected: 2,
		},
		{
			name:     "example4",
			program:  []int{2, 4, 4, 5, 99, 0},
			expected: 2,
		},
		{
			name:     "example5",
			program:  []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			expected: 30,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := gravity.Run(tc.program)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
