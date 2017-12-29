package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jlbribeiro/adventofcode/day19/part2/tubes"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	walker := tubes.NewWalkerFromInput(string(input))
	walker.Walk()
	fmt.Println(walker.Steps())
}
