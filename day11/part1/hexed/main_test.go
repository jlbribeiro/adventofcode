package hexed_test

import (
	"fmt"
	"testing"

	"github.com/jlbribeiro/adventofcode/day11/part1/hexed"
)

var minStepsTests = []struct {
	steps    string
	expected int
}{
	{
		`ne,ne,ne`,
		3,
	},
	{
		`ne,ne,sw,sw`,
		0,
	},
	{
		`ne,ne,s,s`,
		2,
	},
	{
		`se,sw,se,sw,sw`,
		3,
	},
}

func TestMinSteps(t *testing.T) {
	for _, tt := range minStepsTests {
		t.Run(fmt.Sprintf("hexed.MinStepsToStart(steps=%s)", tt.steps), func(t *testing.T) {
			walker := hexed.NewHexWalker()
			walker.WalkFromInput(tt.steps)
			actual := walker.MinStepsToStart()
			if actual != tt.expected {
				t.Errorf("hexed.MinStepsToStart(steps=%s): expected %d, got %d", tt.steps, tt.expected, actual)
			}
		})
	}
}
