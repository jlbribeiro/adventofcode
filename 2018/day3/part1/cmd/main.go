package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2018/day3/part1/fabric"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	var claimInps []string
	for reader.Scan() {
		claimInp := reader.Text()
		claimInps = append(claimInps, claimInp)
	}

	fmt.Println(fabric.OverlapsFromInput(claimInps))
}
