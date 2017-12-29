package maze_test

import (
	"fmt"
	"testing"

	"git.jlbribeiro.com/adventofcode/day5/part2/maze"
)

var walkTests = []struct {
	maze     []int
	expected int
}{
	{
		[]int{0, 3, 0, 1, -3},
		10,
	},
}

func TestWalk(t *testing.T) {
	for i, tt := range walkTests {
		t.Run(fmt.Sprintf("Walk(maze) [%d]", i), func(t *testing.T) {
			actual := maze.Walk(tt.maze)
			if actual != tt.expected {
				t.Errorf("Walk(%v): expected %d, got %d", tt.maze, tt.expected, actual)
			}
		})
	}
}
