package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jlbribeiro/adventofcode/2018/day7/part2/topological"
)

func main() {
	necessaryTime, err := topological.Order(os.Stdin, 5, 60)
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	fmt.Println(necessaryTime)
}
