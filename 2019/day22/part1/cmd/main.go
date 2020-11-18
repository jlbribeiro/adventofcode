package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day22/part1/shuffle"
)

func main() {
	fmt.Println(shuffle.FindCardAfterShuffle(10007, os.Stdin, 2019))
}
