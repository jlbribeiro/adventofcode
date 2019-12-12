package nbody_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day12/part2/nbody"
)

func TestSimulateEnergy(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		steps    int
		expected int
	}{
		{
			name: "example1",
			input: `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`,
			steps:    10,
			expected: 179,
		},
		{
			name: "example2",
			input: `<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>`,
			steps:    100,
			expected: 1940,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.input)
			actual := nbody.SimulateEnergyFromInput(input, tc.steps)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func TestSimulateUntilPreviousState(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		expected int64
	}{
		{
			name: "example1",
			input: `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`,
			expected: 2772,
		},
		{
			name: "example2",
			input: `<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>`,
			expected: 4686774924,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.input)
			actual := nbody.SimulateUntilPreviousStateFromInput(input)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
