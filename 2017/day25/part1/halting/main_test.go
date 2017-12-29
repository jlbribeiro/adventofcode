package halting_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2017/day25/part1/halting"
)

var checksumTests = []struct {
	name      string
	blueprint string
	expected  int
}{
	{
		"example",
		`Begin in state A.
Perform a diagnostic checksum after 6 steps.

In state A:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state B.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the left.
    - Continue with state B.

In state B:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state A.
  If the current value is 1:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state A.`,
		3,
	},
}

func TestChecksum(t *testing.T) {
	for _, tt := range checksumTests {
		t.Run(tt.name, func(t *testing.T) {
			cpu := halting.NewCPUFromBlueprint(tt.blueprint)
			cpu.Run()
			actual := cpu.Checksum()
			if actual != tt.expected {
				t.Errorf("cpu.Checksum(): expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
