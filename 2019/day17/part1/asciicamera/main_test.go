package asciicamera_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day17/part1/asciicamera"
)

func TestSumAlignmentParameters(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "example",
			input: `..#..........
..#..........
#######...###
#.#...#...#.#
#############
..#...#...#..
..#####...^..`,
			expected: 76,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rows := strings.Split(tc.input, "\n")
			input := make([][]rune, len(rows))
			for i, row := range rows {
				input[i] = make([]rune, len(row))
				for j, cell := range row {
					input[i][j] = cell
				}
			}

			actual := asciicamera.SumAlignmentCoordinates(input)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
