package spiral_test

import (
	"testing"

	. "github.com/jlbribeiro/adventofcode/2017/day3/part1/spiral"
)

var manhattanDistanceTests = []struct {
	x1       int
	y1       int
	x2       int
	y2       int
	expected int
}{
	{0, 0, 0, 0, 0},
	{0, 0, 1, 1, 2},
	{0, 0, 1, 0, 1},
	{1, 2, 3, 4, 4},
	{8, 8, 0, 1, 15},
}

func TestManhattanDistance(t *testing.T) {
	for _, tt := range manhattanDistanceTests {
		actual := ManhattanDistance(tt.x1, tt.y1, tt.x2, tt.y2)
		if actual != tt.expected {
			t.Errorf("ManhattanDistance(%d, %d, %d, %d): expected %d, actual %d",
				tt.x1, tt.y1, tt.x2, tt.y2, tt.expected, actual)
		}
	}
}

var spiralCarryDistanceTests = []struct {
	n        int
	expected int
}{
	{1, 0},
	{12, 3},
	{23, 2},
	{1024, 31},
}

func TestSpiralCarryDistance(t *testing.T) {
	for _, tt := range spiralCarryDistanceTests {
		actual := CarryDistance(tt.n)
		if actual != tt.expected {
			t.Errorf("SpiralCarryDistance(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}
