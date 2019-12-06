package orbitmap_test

import "testing"

import "strings"

import "github.com/jlbribeiro/adventofcode/2019/day6/part1/orbitmap"

func TestCountOrbits(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "example",
			input: `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`,
			expected: 42,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.input)
			orbits := orbitmap.ObjectsFromInput(input)
			actual := orbits.CountOrbits()
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
