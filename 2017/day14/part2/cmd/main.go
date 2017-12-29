package main

import (
	"fmt"

	"github.com/jlbribeiro/adventofcode/2017/day14/part2/defragmenter"
)

func main() {
	var key string
	fmt.Scanf("%s", &key)
	fmt.Println(defragmenter.DiskRegionsFromHash(key))
}
