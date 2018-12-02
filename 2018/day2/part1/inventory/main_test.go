package inventory_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2018/day2/part1/inventory"
)

func TestChecksum(t *testing.T) {
	var testCases = []struct {
		name     string
		boxIDs   []string
		expected int
	}{
		{
			name:     "example",
			boxIDs:   []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"},
			expected: 12,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := inventory.Checksum(tc.boxIDs)

			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
