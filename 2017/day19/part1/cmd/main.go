package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jlbribeiro/adventofcode/day19/part1/tubes"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	walker := tubes.NewWalkerFromInput(string(input))
	fmt.Println(walker.Walk())
}
