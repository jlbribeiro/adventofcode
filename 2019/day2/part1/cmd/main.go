package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/jlbribeiro/adventofcode/2019/day2/part1/gravity"
)

func main() {
	inputS, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	programS := strings.Split(strings.TrimSpace(string(inputS)), ",")
	program := make([]int, len(programS))
	for i, nS := range programS {
		n, err := strconv.ParseInt(nS, 10, 32)
		if err != nil {
			panic(err)
		}

		program[i] = int(n)
	}

	program[1], program[2] = 12, 2

	fmt.Println(gravity.Run(program))
}
