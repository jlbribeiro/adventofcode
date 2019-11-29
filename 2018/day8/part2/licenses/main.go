package licenses

import (
	"bufio"
	"io"
	"strconv"
)

func FromTreeInput(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	var input []int
	for scanner.Scan() {
		n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		input = append(input, int(n))
	}

	sum, _ := FromTree(input, 0)
	return sum
}

func FromTree(input []int, index int) (int, int) {
	childNodes := input[index]
	index++
	childValues := make([]int, childNodes, childNodes)

	nMetadataEntries := input[index]
	index++
	metadataEntries := make([]int, nMetadataEntries, nMetadataEntries)

	var childValue int
	for i := range childValues {
		childValue, index = FromTree(input, index)
		childValues[i] = childValue
	}

	for i := range metadataEntries {
		entry := input[index]
		index++

		metadataEntries[i] = entry
	}

	if childNodes == 0 {
		sum := 0
		for _, entry := range metadataEntries {
			sum += entry
		}

		return sum, index
	}

	sum := 0
	for _, entry := range metadataEntries {
		entry--
		if entry < 0 || entry >= len(childValues) {
			continue
		}

		sum += childValues[entry]
	}

	return sum, index
}
