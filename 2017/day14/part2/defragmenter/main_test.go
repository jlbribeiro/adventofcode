package defragmenter_test

import (
	"fmt"
	"testing"

	"github.com/jlbribeiro/adventofcode/2017/day14/part2/defragmenter"
)

var diskRegionsFromHashTests = []struct {
	key      string
	expected int
}{
	{"flqrgnkx", 1242},
}

func TestDiskRegionsFromHash(t *testing.T) {
	for _, tt := range diskRegionsFromHashTests {
		t.Run(fmt.Sprintf("defragmenter.DiskRegionsFromHash(%s)", tt.key), func(t *testing.T) {
			actual := defragmenter.DiskRegionsFromHash(tt.key)
			if actual != tt.expected {
				t.Errorf("defragmenter.DiskRegionsFromHash(%s): expected %d, got %d", tt.key, tt.expected, actual)
			}
		})
	}
}
