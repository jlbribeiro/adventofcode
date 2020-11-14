package doorkey_test

import "testing"

import "github.com/jlbribeiro/adventofcode/2019/day18/part1/doorkey"

import "strings"

func TestShortestPath(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "example1",
			input: `#########
#b.A.@.a#
#########`,
			expected: 8,
		},
		{
			name: "example2",
			input: `########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################`,
			expected: 86,
		},
		{
			name: "example3",
			input: `########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################`,
			expected: 132,
		},
		{
			name: "example4",
			input: `#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################`,
			expected: 136,
		},
		{
			name: "example5",
			input: `########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################`,
			expected: 81,
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
