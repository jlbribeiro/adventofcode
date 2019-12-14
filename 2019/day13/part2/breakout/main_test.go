package breakout_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day13/part2/breakout"
)

func TestBlockTilesCount(t *testing.T) {
	var testCases = []struct {
		name     string
		level    []int64
		expected int
	}{
		{
			name:     "custom_example",
			level:    []int64{0, 0, 0, 0, 1, 1, 0, 2, 2, 0, 3, 2, 0, 4, 2},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := breakout.BlockTilesCount(tc.level)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
