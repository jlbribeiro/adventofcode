package fabric_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2018/day3/part2/fabric"
)

func TestOverlaps(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "example",
			input: `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`,
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			claimInps := strings.Split(tc.input, "\n")
			actual := fabric.OverlapsFromInput(claimInps)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
