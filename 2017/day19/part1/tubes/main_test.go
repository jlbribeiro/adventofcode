package tubes_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/2017/day19/part1/tubes"
)

var walkTests = []struct {
	name     string
	input    string
	expected string
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
		"ABCDEF",
	},
}

func TestWalk(t *testing.T) {
	for _, tt := range walkTests {
		t.Run(tt.name, func(t *testing.T) {
			walker := tubes.NewWalkerFromInput(tt.input)

			actual := walker.Walk()
			if actual != tt.expected {
				t.Errorf("tubes.Walk(%s): expected %s, got %s", tt.name, tt.expected, actual)
			}
		})
	}
}
