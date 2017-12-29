package coprocessor_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2017/day23/part2/coprocessor"
)

var hValueTests = []struct {
	name     string
	program  []string
	expected int
}{
	{
		"made_up",
		[]string{
			"set a 3",
			"set h 3",
			"mul h a",
			"mod h 5",
		},
		4,
	},
}

func TestHValue(t *testing.T) {
	for _, tt := range hValueTests {
		t.Run(tt.name, func(t *testing.T) {
			cp := coprocessor.NewCoprocessor()
			cp.Run(tt.program)
			actual := cp.GetValue("h")
			if actual != tt.expected {
				t.Errorf("coprocessor.MulCalls(%s): expected %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}
