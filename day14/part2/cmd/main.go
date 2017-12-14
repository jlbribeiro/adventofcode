package main

import (
	"fmt"

	"github.com/jlbribeiro/adventofcode/day14/part2/defragmenter"
)

func main() {
	var key string
	fmt.Scanf("%s", &key)
	fmt.Println(defragmenter.DiskRegionsFromHash(key))
}
