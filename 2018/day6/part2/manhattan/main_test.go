package manhattan_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2018/day6/part2/manhattan"
)

func TestRegionSizeOfMaxDistance(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		distance int
		expected int
	}{
		{
			name: "example",
			input: `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`,
			distance: 32,
			expected: 16,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := manhattan.RegionSizeOfMaxDistance(strings.NewReader(tc.input), tc.distance)
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
