package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day22/part2/shuffle"
)

func main() {
	var nCards int64 = 119315717514047
	var nTimes int64 = 101741582076661
	var index int64 = 2020
	fmt.Println(shuffle.GetAfterShuffle(nCards, os.Stdin, nTimes, index))
}
