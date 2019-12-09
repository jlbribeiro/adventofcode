package thrusters_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day7/part1/thrusters"
)

func TestAmplifyFromInput(t *testing.T) {
	var testCases = []struct {
		name         string
		inputProgram string
		expected     int
	}{
		{
			name:         "example1",
			inputProgram: "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0",
			expected:     43210,
		},
		{
			name:         "example2",
			inputProgram: "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0",
			expected:     54321,
		},
		{
			name:         "example3",
			inputProgram: "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0",
			expected:     65210,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.inputProgram)
			actual := thrusters.AmplifyFromInput(input, 5)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}