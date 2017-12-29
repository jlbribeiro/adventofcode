package firewall_test

import (
	"fmt"
	"testing"

	"github.com/jlbribeiro/adventofcode/2017/day13/part2/firewall"
)

type testLayer struct {
	depth         int
	scanningRange int
}

var undetectedWalkDelayTests = []struct {
	layers   []testLayer
	expected int
}{
	{
		layers: []testLayer{
			{
				depth:         0,
				scanningRange: 3,
			},
			{
				depth:         1,
				scanningRange: 2,
			},
			{
				depth:         4,
				scanningRange: 4,
			},
			{
				depth:         6,
				scanningRange: 4,
			},
		},
		expected: 10,
	},
}

func TestUndetectedWalkDelay(t *testing.T) {
	for i, tt := range undetectedWalkDelayTests {
		t.Run(fmt.Sprintf("firewall.UndetectedWalkDelay(test_i=%d)", i), func(t *testing.T) {
			fw := firewall.NewFirewall()
			for _, layer := range tt.layers {
				fw.AddLayer(firewall.NewLayer(layer.depth, layer.scanningRange))
			}

			actual := fw.UndetectedWalkDelay()
			if actual != tt.expected {
				t.Errorf("firewall.UndetectedWalkDelay(test_i=%d): expected %d, got %d", i, tt.expected, actual)
			}
		})
	}
}
