package main

import (
	"flag"
	"log"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day25/part1/cryo"
)

func main() {
	var inputFilepath string
	flag.StringVar(&inputFilepath, "input", "", "input filepath")
	flag.Parse()

	if inputFilepath == "" {
		flag.Usage()
		os.Exit(-1)
	}

	input, err := os.Open(inputFilepath)
	if err != nil {
		log.Fatalf("failed to open input file: %s", inputFilepath)
	}

	cryo.WalkFromInput(input, os.Stdin)
}
