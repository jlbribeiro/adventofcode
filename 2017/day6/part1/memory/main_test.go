package memory_test

import (
	"fmt"
	"testing"

	"git.jlbribeiro.com/adventofcode/2017/day6/part1/memory"
)

var rebalanceLoopTests = []struct {
	banks    []int
	expected int
}{
	{[]int{0, 2, 7, 0}, 5},
}

func TestRebalanceLoop(t *testing.T) {
	for _, tt := range rebalanceLoopTests {
		t.Run(fmt.Sprintf("RebalanceLoop(%v)", tt.banks), func(t *testing.T) {
			actual := memory.RebalanceLoop(tt.banks)
			if actual != tt.expected {
				t.Errorf("RebalanceLoop(%v): expected %d, got %d", tt.banks, tt.expected, actual)
			}
		})
	}
}
