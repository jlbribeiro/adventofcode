package doorkey_test

import "testing"

import "github.com/jlbribeiro/adventofcode/2019/day18/part2/doorkey"

import "strings"

func TestShortestPath(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "example1",
			input: `#######
#a.#Cd#
##@#@##
#######
##@#@##
#cB#Ab#
#######`,
			expected: 8,
		},
		{
			name: "example2",
			input: `###############
#d.ABC.#.....a#
######@#@######
###############
######@#@######
#b.....#.....c#
###############`,
			expected: 24,
		},
		{
			name: "example3",
			input: `#############
#DcBa.#.GhKl#
#.###@#@#I###
#e#d#####j#k#
###C#@#@###J#
#fEbA.#.FgHi#
#############`,
			expected: 32,
		},
		{
			name: "example4",
			input: `#############
#g#f.D#..h#l#
#F###e#E###.#
#dCba@#@BcIJ#
#############
#nK.L@#@G...#
#M###N#H###.#
#o#m..#i#jk.#
#############`,
			expected: 72,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.input)
			actual := doorkey.ShortestPathFromInput(input)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
