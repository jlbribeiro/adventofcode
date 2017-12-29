package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/day10/part2/knothash"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	input, err := scanner.ReadBytes('\n')
	if err != nil {
		panic(err)
	}

	// Skip newline
	input = input[:len(input)-1]

	hash := knothash.NewHash(256)
	digest := hash.Digest(input)
	fmt.Println(digest)
}
