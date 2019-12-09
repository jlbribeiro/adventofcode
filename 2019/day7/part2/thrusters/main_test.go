package thrusters_test

import (
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/2019/day7/part2/thrusters"
)

func TestAmplifyFromInput(t *testing.T) {
	var testCases = []struct {
		name         string
		inputProgram string
		expected     int
	}{
		{
			name:         "example1",
			inputProgram: "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5",
			expected:     139629729,
		},
		{
			name:         "example2",
			inputProgram: "3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10",
			expected:     18216,
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
