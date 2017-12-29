package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2017/day12/part1/digitalplumber"
)

func main() {
	pn := digitalplumber.NewProgramNetworkFromReader(os.Stdin)
	fmt.Println(pn.NConnectionsOf(0))
}
