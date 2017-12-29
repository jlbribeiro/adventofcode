package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/day12/part2/digitalplumber"
)

func main() {
	pn := digitalplumber.NewProgramNetworkFromReader(os.Stdin)
	pn.Flood()
	fmt.Println(pn.NGroups())
}
