package inventory_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2018/day2/part2/inventory"
)

func TestFindCorrectBoxes(t *testing.T) {
	var testCases = []struct {
		name     string
		boxIDs   []string
		expected string
	}{
		{
			name: "example",
			boxIDs: []string{
				"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz",
			},
			expected: "fgij",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := inventory.FindCorrectBoxes(tc.boxIDs)

			if actual != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, actual)
			}
		})
	}
}
