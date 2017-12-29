package defragmenter

import (
	"fmt"

	"github.com/jlbribeiro/adventofcode/day14/part1/knothash"
)

func DiskUsageFromHash(key string) int {
	diskUsage := 0

	for i := 0; i < 128; i++ {
		rowKey := fmt.Sprintf("%s-%d", key, i)

		hash := knothash.NewHash(256)
		hash.Digest([]byte(rowKey))
		rowHash := hash.DenseHash()

		for _, b := range rowHash {
			for j := 0; j < 8; j++ {
				diskUsage += int(b & 0x1)
				b = b >> 1
			}
		}
	}

	return diskUsage
}
