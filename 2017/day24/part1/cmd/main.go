package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2017/day24/part1/bridges"
)

func main() {
	pieces := []string(nil)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		piece := scanner.Text()
		pieces = append(pieces, piece)
	}

	bb := bridges.NewBridgeBuilder(pieces)
	fmt.Println(bb.StrongestBridge())
}
