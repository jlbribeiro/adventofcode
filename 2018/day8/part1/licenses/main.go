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
	sum := 0

	childNodes := input[index]
	index++

	metadataEntries := input[index]
	index++

	var subsum int
	for i := 0; i < childNodes; i++ {
		subsum, index = FromTree(input, index)
		sum += subsum
	}

	for i := 0; i < metadataEntries; i++ {
		entry := input[index]
		sum += entry
		index++
	}

	return sum, index
}
