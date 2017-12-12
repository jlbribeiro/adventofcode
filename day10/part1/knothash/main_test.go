package knothash_test

import (
	"fmt"
	"testing"

	"github.com/jlbribeiro/adventofcode/day10/part1/knothash"
)

var hashTests = []struct {
	n        int
	lengths  []int
	expected int
}{
	{
		5,
		[]int{3, 4, 1, 5},
		12,
	},
}

func TestHash(t *testing.T) {
	for _, tt := range hashTests {
		t.Run(fmt.Sprintf("knothash.Hash(%d, %v)", tt.n, tt.lengths), func(t *testing.T) {
			actual := knothash.Hash(tt.n, tt.lengths)
			if actual != tt.expected {
				t.Errorf("knothash.Hash(%d, %v): expected %d, got %d", tt.n, tt.lengths, tt.expected, actual)
			}
		})
	}
}
