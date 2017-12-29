package circus_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jlbribeiro/adventofcode/day7/part1/circus"
)

var programCircusTests = []struct {
	input    string
	expected string
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
		"tknk",
	},
}

func TestProgramCircus(t *testing.T) {
	for _, tt := range programCircusTests {
		t.Run(fmt.Sprintf("TestProgramCircus(expected = %s)", tt.expected), func(t *testing.T) {
			lines := strings.Split(tt.input, "\n")
			tower := circus.NewTower()
			for _, programYell := range lines {
				tower.RegisterProgram(circus.NewProgramFromYell(programYell))
			}

			program := tower.FindBottomProgram()
			if program == nil || program.Name != tt.expected {
				t.Errorf("")
			}
		})
	}
}
