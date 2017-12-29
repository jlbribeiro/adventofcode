package bridges_test

import (
	"testing"

	"github.com/jlbribeiro/adventofcode/day24/part1/bridges"
)

var strongestBridgeTests = []struct {
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
		31,
	},
}

func TestStrongestBridge(t *testing.T) {
	for _, tt := range strongestBridgeTests {
		t.Run(tt.name, func(t *testing.T) {
			bridgeBuilder := bridges.NewBridgeBuilder(tt.components)
			actual := bridgeBuilder.StrongestBridge()
			if actual != tt.expected {
				t.Errorf("bridgeBuilder.StrongestBridge(): expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
