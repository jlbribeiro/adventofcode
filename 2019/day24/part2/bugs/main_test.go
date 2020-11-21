package bugs_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day24/part2/bugs"
)

func TestBugsAfterNGenerations(t *testing.T) {
	var testCases = []struct {
		name         string
		input        string
		nGenerations int
		expected     int
	}{
		{
			name: "example1",
			input: `....#
#..#.
#..##
..#..
#....`,
			nGenerations: 10,
			expected:     99,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.input)
			actual := bugs.BugsAfterNGenerationsFromInput(input, tc.nGenerations)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
