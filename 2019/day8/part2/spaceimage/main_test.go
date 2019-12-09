package spaceimage_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day8/part2/spaceimage"
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

func TestRenderFromInput(t *testing.T) {
	var testCases = []struct {
		name     string
		width    int
		height   int
		input    string
		expected [][]int
	}{
		{
			name:   "example",
			width:  2,
			height: 2,
			input:  `0222112222120000`,
			expected: [][]int{
				[]int{0, 1},
				[]int{1, 0},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.input)
			actual := spaceimage.RenderFromInput(input, tc.width, tc.height)
			if len(actual) != len(tc.expected) {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}

			for i := 0; i < len(tc.expected); i++ {
				if len(actual[i]) != len(tc.expected[i]) {
					t.Fatalf("expected %v, got %v", tc.expected, actual)
				}

				for j := 0; j < len(tc.expected[0]); j++ {
					if actual[i][j] != tc.expected[i][j] {
						t.Errorf("expected %v, got %v", tc.expected, actual)
					}
				}
			}
		})
	}
}
