package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jlbribeiro/adventofcode/2017/day25/part1/halting"
)

func main() {
	inputBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	input := string(inputBytes)

	cpu := halting.NewCPUFromBlueprint(input)
	cpu.Run()
	fmt.Println(cpu.Checksum())
}
