package springdroid

import (
	"fmt"
	"io"
	"strings"

	"github.com/jlbribeiro/adventofcode/2019/day21/part1/intcode"
)

func HullDamageFromInput(input io.Reader) int {
	program := intcode.ProgramFromInput(input)
	cpu := intcode.NewCPU(program)

	// J = (!A ^ D) v (!B ^ D) v (!C ^ D)
	//   = (!A v !B v !C) ^ D
	//   = !(A ^ B ^ C) ^ D
	instructions := []rune(strings.Join([]string{
		"NOT A T", // T = !A
		"NOT T T", // T = A

		"AND B T", // T = A ^ B

		"AND C T", // T = (A ^ B) ^ C

		"NOT T J", // J = !(A ^ B ^ C)

		"AND D J", // J = !(A ^ B ^ C) ^ D

		"WALK",
		"\n",
	}, "\n"))
	instructionsI64 := make([]int64, len(instructions))
	for i := range instructions {
		instructionsI64[i] = int64(instructions[i])
	}

	outputI64, _ := cpu.Exec(instructionsI64)

	for _, ch := range outputI64 {
		if ch >= 0 && ch <= 255 {
			fmt.Print(string(rune(ch)))
			continue
		}

		return int(ch)
	}

	return -1
}
