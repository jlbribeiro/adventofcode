package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jlbribeiro/adventofcode/2018/day1/part2/freq"
)

func main() {
	var deltas []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		deltaIn := scanner.Text()
		delta, err := strconv.ParseInt(deltaIn, 10, 64)
		if err != nil {
			log.Fatalf("invalid delta: %s", deltaIn)
		}

		deltas = append(deltas, int(delta))
	}

	fmt.Println(freq.Analyse(deltas))
}
