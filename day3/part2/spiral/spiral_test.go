package spiral_test

import (
	"testing"

	"git.jlbribeiro.com/adventofcode/day3/part2/spiral"
)

var largerSpiralSumTests = []struct {
	n        int
	expected int
}{
	{25, 26},
	{120, 122},
	{748, 806},
}

func TestLargerSpiralSum(t *testing.T) {
	for _, tt := range largerSpiralSumTests {
		actual := spiral.GetSumLargerThan(tt.n)
		if actual != tt.expected {
			t.Errorf("LargerSpiralSum(%d): expected %d, got %d", tt.n, tt.expected, actual)
		}
	}
}
