package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2018/day7/part1/topological"
)

func main() {
	fmt.Println(topological.Order(os.Stdin))
}
