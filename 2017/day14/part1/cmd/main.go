package main

import (
	"fmt"

	"github.com/jlbribeiro/adventofcode/day14/part1/defragmenter"
)

func main() {
	var key string
	fmt.Scanf("%s", &key)
	fmt.Println(defragmenter.DiskUsageFromHash(key))
}
