package defragmenter_test

import (
	"fmt"
	"testing"

	"github.com/jlbribeiro/adventofcode/day14/part1/defragmenter"
)

var diskUsageFromHashTests = []struct {
	key      string
	expected int
}{
	{"flqrgnkx", 8108},
}

func TestDiskUsageFromHash(t *testing.T) {
	for _, tt := range diskUsageFromHashTests {
		t.Run(fmt.Sprintf("defragmenter.DiskUsageFromHash(%s)", tt.key), func(t *testing.T) {
			actual := defragmenter.DiskUsageFromHash(tt.key)
			if actual != tt.expected {
				t.Errorf("defragmenter.DiskUsageFromHash(%s): expected %d, got %d", tt.key, tt.expected, actual)
			}
		})
	}
}
