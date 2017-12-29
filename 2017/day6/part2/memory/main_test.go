package memory_test

import (
	"fmt"
	"testing"

	"git.jlbribeiro.com/adventofcode/2017/day6/part2/memory"
)

var rebalanceLoopTests = []struct {
	banks    []int
	expected int
}{
	{[]int{0, 2, 7, 0}, 4},
}

func TestRebalanceRepeatLoop(t *testing.T) {
	for _, tt := range rebalanceLoopTests {
		t.Run(fmt.Sprintf("RebalanceLoop(%v)", tt.banks), func(t *testing.T) {
			stateHistory := memory.NewStateHistory()
			actual := memory.RebalanceRepeatLoop(stateHistory, tt.banks)
			if actual != tt.expected {
				t.Errorf("RebalanceLoop(%v): expected %d, got %d", tt.banks, tt.expected, actual)
			}
		})
	}
}
