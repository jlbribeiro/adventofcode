package circus_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/day7/part2/circus"
)

var programCircusTests = []struct {
	input    string
	expected int
}{
	{
		`pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`,
		60,
	},
}

func TestProgramCircus(t *testing.T) {
	for _, tt := range programCircusTests {
		t.Run(fmt.Sprintf("TestProgramCircus(expected = %d)", tt.expected), func(t *testing.T) {
			lines := strings.Split(tt.input, "\n")
			tower := circus.NewTower()
			for _, programYell := range lines {
				tower.RegisterProgram(circus.NewProgramFromYell(programYell))
			}

			actual := tower.FindWrongWeightProgramIdealWeight()
			if actual != tt.expected {
				t.Errorf("TestProgramCircus(): expected %d, got %d", tt.expected, actual)
			}
		})
	}
}

var differenceTests = []struct {
	list         []int
	expectedDiff int
	expectedInd  int
}{
	// Less than 3 elements: no difference can be calculated.
	{[]int{}, 0, -1},
	{[]int{1}, 0, -1},
	{[]int{1, 2}, 0, -1},

	// Edge cases: first and second element.
	// The difference's signal is important.
	{[]int{0, 0, 0}, 0, -1},
	{[]int{0, 7, 0}, -7, 1},
	{[]int{2, 0, 0}, -2, 0},
	{[]int{7, 2, 7}, 5, 1},
	{[]int{2, 7, 7}, 5, 0},

	{[]int{0, 0, 0, 1}, -1, 3},
	{[]int{3, 3, 3, 3}, 0, -1},
}

func TestDifference(t *testing.T) {
	for _, tt := range differenceTests {
		t.Run(fmt.Sprintf("TestDifference(%v)", tt.list), func(t *testing.T) {
			actualDiff, actualInd := circus.Difference(tt.list)
			if actualDiff != tt.expectedDiff || actualInd != tt.expectedInd {
				t.Errorf("TestDifference(%v): expected (%d, %d), got (%d, %d)",
					tt.list, tt.expectedDiff, tt.expectedInd,
					actualDiff, actualInd)
			}
		})
	}
}
