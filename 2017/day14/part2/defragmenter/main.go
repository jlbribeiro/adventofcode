package defragmenter

import (
	"fmt"

	"github.com/jlbribeiro/adventofcode/day14/part2/knothash"
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

func DiskRegionsFromHash(key string) int {
	disk := make([][]int, 128)
	for i := range disk {
		disk[i] = make([]int, 128)
	}

	for rowID, row := range disk {
		rowKey := fmt.Sprintf("%s-%d", key, rowID)

		hash := knothash.NewHash(256)
		hash.Digest([]byte(rowKey))
		rowHash := hash.DenseHash()

		for blockID, byt := range rowHash {
			for bit := 7; bit >= 0; bit-- {
				row[blockID*8+bit] = int(byt & 0x1)
				byt = byt >> 1
			}
		}
	}

	// Start regionID as 1 since it increments before flooding, which means the
	// first regionID will be 2 (since 0 means unused and 1 means used but
	// region unidentified).
	regionID := 1

	for i := range disk {
		for j := range disk[i] {
			// Disk is used but region is unidentified.
			// Since identifyDiskRegion only returns when a full region has
			// been identified, we're sure we can increment the regionID each
			// time used disk space without an identified region is found.
			if disk[i][j] == 1 {
				regionID++
				identifyDiskRegion(disk, i, j, regionID)
			}
		}
	}

	// regionID will have +1 because we need to start the count at 2.
	return regionID - 1
}

func identifyDiskRegion(disk [][]int, rowID int, colID int, regionID int) {
	// Check limits.
	if rowID < 0 || rowID >= len(disk) || colID < 0 || colID >= len(disk[rowID]) {
		return
	}

	// Unused or already identified.
	if disk[rowID][colID] == 0 || disk[rowID][colID] > 1 {
		return
	}

	// Unmarked used disk. Mark it as regionID.
	disk[rowID][colID] = regionID

	identifyDiskRegion(disk, rowID-1, colID, regionID)
	identifyDiskRegion(disk, rowID, colID-1, regionID)
	identifyDiskRegion(disk, rowID, colID+1, regionID)
	identifyDiskRegion(disk, rowID+1, colID, regionID)
}
