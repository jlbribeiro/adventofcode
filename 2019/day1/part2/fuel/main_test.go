package fuel_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day1/part2/fuel"
)

func TestModuleRequirements(t *testing.T) {
	var testCases = []struct {
		name     string
		mass     int
		expected int
	}{
		{
			name:     "example1",
			mass:     14,
			expected: 2,
		},
		{
			name:     "example2",
			mass:     1969,
			expected: 966,
		},
		{
			name:     "example3",
			mass:     100756,
			expected: 50346,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cache := fuel.NewFuelCache()
			actual := cache.ModuleRequirements(tc.mass)
			if actual != tc.expected {
				t.Errorf("expected %v, got %v\n", tc.expected, actual)
			}
		})
	}
}
