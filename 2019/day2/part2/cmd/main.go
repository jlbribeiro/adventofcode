package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/jlbribeiro/adventofcode/2019/day2/part2/gravity"
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

	targetOutput := 19690720
	memory := make([]int, len(initialMemory), len(initialMemory))
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(memory, initialMemory)

			memory[1], memory[2] = noun, verb
			if gravity.Run(memory) == targetOutput {
				fmt.Println(100*noun + verb)
				return
			}
		}
	}

	fmt.Printf("Noun and Verb not found for target output %d\n", targetOutput)
}
