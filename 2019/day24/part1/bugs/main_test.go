package bugs_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day24/part1/bugs"
)

func TestBugsRepeatedBiodiversity(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "example1",
			input: `....#
#..#.
#..##
..#..
#....`,
			expected: 2129920,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.input)
			actual := bugs.RepeatedBiodiversityFromInput(input)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
