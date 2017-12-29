package bridges_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/day24/part2/bridges"
)

var longestStrongestBridgeTests = []struct {
	name       string
	components []string
	expected   int
}{
	{
		"example",
		[]string{
			"0/2",
			"2/2",
			"2/3",
			"3/4",
			"3/5",
			"0/1",
			"10/1",
			"9/10",
		},
		19,
	},
}

func TestStrongestBridge(t *testing.T) {
	for _, tt := range longestStrongestBridgeTests {
		t.Run(tt.name, func(t *testing.T) {
			bridgeBuilder := bridges.NewBridgeBuilder(tt.components)
			actual := bridgeBuilder.LongestStrongestBridge()
			if actual != tt.expected {
				t.Errorf("bridgeBuilder.StrongestBridge(): expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
