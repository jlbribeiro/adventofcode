package registers_test

import (
	"fmt"
	"testing"

	"github.com/jlbribeiro/adventofcode/2017/day8/part2/registers"
)

var cpuGetLargestRegisterValueTests = []struct {
	input    string
	expected int
}{
	{
		`b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`,
		10,
	},
}

func TestCPUGetLargestRegisterValueEverTests(t *testing.T) {
	for i, tt := range cpuGetLargestRegisterValueTests {
		t.Run(fmt.Sprintf("CPU.GetLargestRegisterValue(testID = %d)", i), func(t *testing.T) {
			cpu := registers.NewCPUFromProgramInput(tt.input)
			cpu.RunProgram()
			actual := cpu.GetLargestRegisterValueEver()
			if actual != tt.expected {
				t.Errorf("CPU.GetLargestRegisterValue(testID = %d): expected %d, got %d", i, tt.expected, actual)
			}
		})
	}
}
