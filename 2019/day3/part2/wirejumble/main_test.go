package wirejumble_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day3/part2/wirejumble"
)

func TestIntersectionDistance(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "example1",
			input: `R8,U5,L5,D3
U7,R6,D4,L4`,
			expected: 30,
		},
		{
			name: "example2",
			input: `R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`,
			expected: 610,
		},
		{
			name: "example3",
			input: `R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`,
			expected: 410,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.input)
			actual := wirejumble.IntersectionDistanceFromInput(input)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
