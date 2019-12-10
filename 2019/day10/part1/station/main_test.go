package station_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day10/part1/station"
)

func TestLoSAt(t *testing.T) {
	grid1 := `.#..#
.....
#####
....#
...##`

	grid5 := `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`

	var testCases = []struct {
		name     string
		inputG   string
		inputX   int
		inputY   int
		expected int
	}{
		{
			name:     "example1/1",
			inputG:   grid1,
			inputX:   1,
			inputY:   0,
			expected: 7,
		},
		{
			name:     "example1/2",
			inputG:   grid1,
			inputX:   4,
			inputY:   0,
			expected: 7,
		},
		{
			name:     "example1/3",
			inputG:   grid1,
			inputX:   0,
			inputY:   2,
			expected: 6,
		},
		{
			name:     "example1/7",
			inputG:   grid1,
			inputX:   4,
			inputY:   2,
			expected: 5,
		},
		{
			name:     "example1/9",
			inputG:   grid1,
			inputX:   3,
			inputY:   4,
			expected: 8,
		},
		{
			name:     "example5/best",
			inputG:   grid5,
			inputX:   11,
			inputY:   13,
			expected: 210,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inputG := strings.NewReader(tc.inputG)
			grid := station.AsteroidsGridFromInput(inputG)
			actual := station.LoSAt(grid, tc.inputX, tc.inputY)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}

}

func TestBestLoSLocation(t *testing.T) {
	var testCases = []struct {
		name        string
		input       string
		expectedLoS int
		expectedX   int
		expectedY   int
	}{
		{
			name: "example1",
			input: `.#..#
.....
#####
....#
...##`,
			expectedLoS: 8,
			expectedX:   3,
			expectedY:   4,
		},
		{
			name: "example2",
			input: `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`,
			expectedLoS: 33,
			expectedX:   5,
			expectedY:   8,
		},
		{
			name: "example3",
			input: `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`,
			expectedLoS: 35,
			expectedX:   1,
			expectedY:   2,
		},
		{
			name: "example4",
			input: `.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`,
			expectedLoS: 41,
			expectedX:   6,
			expectedY:   3,
		},
		{
			name: "example5",
			input: `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`,
			expectedLoS: 210,
			expectedX:   11,
			expectedY:   13,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.input)
			grid := station.AsteroidsGridFromInput(input)
			actualLoS, actualX, actualY := station.BestLoSLocation(grid)

			if actualLoS != tc.expectedLoS || actualX != tc.expectedX || actualY != tc.expectedY {
				t.Errorf("expected (los=%d, (%d, %d)), got (los=%d, (%d, %d))", tc.expectedLoS, tc.expectedX, tc.expectedY, actualLoS, actualX, actualY)
			}
		})
	}
}
