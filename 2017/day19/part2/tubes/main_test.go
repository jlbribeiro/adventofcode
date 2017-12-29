package tubes_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/day19/part2/tubes"
)

var walkStepsTests = []struct {
	name     string
	input    string
	expected int
}{
	{
		"example",
		`     |          
     |  +--+    
     A  |  C    
 F---|----E|--+ 
     |  |  |  D 
     +B-+  +--+ 
`,
		38,
	},
}

func TestWalkSteps(t *testing.T) {
	for _, tt := range walkStepsTests {
		t.Run(tt.name, func(t *testing.T) {
			walker := tubes.NewWalkerFromInput(tt.input)
			walker.Walk()
			actual := walker.Steps()
			if actual != tt.expected {
				t.Errorf("tubes.Walk(%s): expected %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}
