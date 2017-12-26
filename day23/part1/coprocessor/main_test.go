package coprocessor_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/day23/part1/coprocessor"
)

var mulCallTests = []struct {
	name     string
	program  []string
	expected int
}{
	{
		"made_up",
		[]string{
			"set a 3",
			"sub b 4",
			"mul d a",
		},
		1,
	},
}

func TestMulCalls(t *testing.T) {
	for _, tt := range mulCallTests {
		t.Run(tt.name, func(t *testing.T) {
			cp := coprocessor.NewCoprocessor()
			cp.Run(tt.program)
			actual := cp.MulCalls()
			if actual != tt.expected {
				t.Errorf("coprocessor.MulCalls(%s): expected %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}
