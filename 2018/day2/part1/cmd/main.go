package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2018/day2/part1/inventory"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var boxIDs []string
	for scanner.Scan() {
		boxID := scanner.Text()
		boxIDs = append(boxIDs, boxID)
	}

	fmt.Println(inventory.Checksum(boxIDs))
}
