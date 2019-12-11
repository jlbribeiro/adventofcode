package station_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day10/part2/station"
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

func TestLaserSweep(t *testing.T) {
	gridSmall := `.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....#...###..
..#.#.....#....##`

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
		name      string
		inputG    string
		inputX    int
		inputY    int
		inputNth  int
		expectedX int
		expectedY int
	}{
		{
			name:      "example_small/1",
			inputG:    gridSmall,
			inputX:    8,
			inputY:    3,
			inputNth:  1,
			expectedX: 8,
			expectedY: 1,
		},
		{
			name:      "example_small/2",
			inputG:    gridSmall,
			inputX:    8,
			inputY:    3,
			inputNth:  2,
			expectedX: 9,
			expectedY: 0,
		},
		{
			name:      "example_small/3",
			inputG:    gridSmall,
			inputX:    8,
			inputY:    3,
			inputNth:  3,
			expectedX: 9,
			expectedY: 1,
		},
		{
			name:      "example_small/4",
			inputG:    gridSmall,
			inputX:    8,
			inputY:    3,
			inputNth:  4,
			expectedX: 10,
			expectedY: 0,
		},
		{
			name:      "example_small/5",
			inputG:    gridSmall,
			inputX:    8,
			inputY:    3,
			inputNth:  5,
			expectedX: 9,
			expectedY: 2,
		},
		{
			name:      "example5/1",
			inputG:    grid5,
			inputX:    11,
			inputY:    13,
			inputNth:  1,
			expectedX: 11,
			expectedY: 12,
		},
		{
			name:      "example5/2",
			inputG:    grid5,
			inputX:    11,
			inputY:    13,
			inputNth:  2,
			expectedX: 12,
			expectedY: 1,
		},
		{
			name:      "example5/3",
			inputG:    grid5,
			inputX:    11,
			inputY:    13,
			inputNth:  3,
			expectedX: 12,
			expectedY: 2,
		},
		{
			name:      "example5/10",
			inputG:    grid5,
			inputX:    11,
			inputY:    13,
			inputNth:  10,
			expectedX: 12,
			expectedY: 8,
		},
		{
			name:      "example5/20",
			inputG:    grid5,
			inputX:    11,
			inputY:    13,
			inputNth:  20,
			expectedX: 16,
			expectedY: 0,
		},
		{
			name:      "example5/50",
			inputG:    grid5,
			inputX:    11,
			inputY:    13,
			inputNth:  50,
			expectedX: 16,
			expectedY: 9,
		},
		{
			name:      "example5/100",
			inputG:    grid5,
			inputX:    11,
			inputY:    13,
			inputNth:  100,
			expectedX: 10,
			expectedY: 16,
		},
		{
			name:      "example5/199",
			inputG:    grid5,
			inputX:    11,
			inputY:    13,
			inputNth:  199,
			expectedX: 9,
			expectedY: 6,
		},
		{
			name:      "example5/200",
			inputG:    grid5,
			inputX:    11,
			inputY:    13,
			inputNth:  200,
			expectedX: 8,
			expectedY: 2,
		},
		{
			name:      "example5/201",
			inputG:    grid5,
			inputX:    11,
			inputY:    13,
			inputNth:  201,
			expectedX: 10,
			expectedY: 9,
		},
		{
			name:      "example5/299",
			inputG:    grid5,
			inputX:    11,
			inputY:    13,
			inputNth:  299,
			expectedX: 11,
			expectedY: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.inputG)
			grid := station.AsteroidsGridFromInput(input)
			actualX, actualY := station.LaserSweep(grid, tc.inputX, tc.inputY, tc.inputNth)
			if actualX != tc.expectedX || actualY != tc.expectedY {
				t.Errorf("expected (%d,%d), got (%d,%d)", tc.expectedX, tc.expectedY, actualX, actualY)
			}
		})
	}
}
