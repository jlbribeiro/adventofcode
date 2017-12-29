package spinlock_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2017/day17/part1/spinlock"
)

var spinlockRunTT = []struct {
	name        string
	step        int
	nIterations int
	expected    int
}{
	{"provided example", 3, 2017, 638},
}

func TestSpinlockRun(t *testing.T) {
	for _, tc := range spinlockRunTT {
		t.Run(tc.name, func(t *testing.T) {
			spinlock := spinlock.NewSpinlock(tc.step)
			actual := spinlock.Run(tc.nIterations)
			if actual != tc.expected {
				t.Errorf("spinlock.Run(): expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
