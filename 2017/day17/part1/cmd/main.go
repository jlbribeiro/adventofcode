package main

import (
	"fmt"

	"github.com/jlbribeiro/adventofcode/2017/day17/part1/spinlock"
)

func main() {
	spinlock := spinlock.NewSpinlock(335)
	fmt.Println(spinlock.Run(2017))
}
