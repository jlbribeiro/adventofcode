package manhattan_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2018/day6/part1/manhattan"
)

func TestLargestArea(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
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
			expected: 17,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := manhattan.LargestArea(strings.NewReader(tc.input))
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
