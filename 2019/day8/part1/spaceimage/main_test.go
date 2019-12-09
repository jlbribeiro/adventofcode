package spaceimage_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day8/part1/spaceimage"
)

func TestChecksumFromInput(t *testing.T) {
	var testCases = []struct {
		name     string
		width    int
		height   int
		input    string
		expected int
	}{
		{
			name:     "example",
			width:    3,
			height:   2,
			input:    `123456789012`,
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.input)
			actual := spaceimage.ChecksumFromInput(input, tc.width, tc.height)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
