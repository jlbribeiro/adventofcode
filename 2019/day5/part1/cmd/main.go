package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/jlbribeiro/adventofcode/2019/day5/part1/thermal"
)

func main() {
	inputS, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	initialMemoryS := strings.Split(strings.TrimSpace(string(inputS)), ",")
	initialMemory := make([]int, len(initialMemoryS))
	for i, nS := range initialMemoryS {
		n, err := strconv.ParseInt(nS, 10, 32)
		if err != nil {
			panic(err)
		}

		initialMemory[i] = int(n)
	}

	memory := make([]int, len(initialMemory), len(initialMemory))
	copy(memory, initialMemory)
	output := thermal.Run(memory, 1)
	fmt.Println(output)
}
