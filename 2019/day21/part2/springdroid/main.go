package springdroid

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/jlbribeiro/adventofcode/2019/day21/part1/intcode"
)

func HullDamageFromInput(input io.Reader) int {
	debugPlayMovie := false

	program := intcode.ProgramFromInput(input)
	cpu := intcode.NewCPU(program)

	// Now that we see a bit further ahead, we know we shouldn't jump
	// just because C is not ground (and D is), as this leads
	// to the scenario below. I made variations of the scenario to justify the
	// modification.
	//
	// Left: the actual result of merely executing the part 1's program;
	// we jump too soon, and fail to jump the second time.
	// That's because the second time, D is not land
	// (which correspond's to the original jump's H sensor).
	//
	// Center: what would happen even if we modified our program to "magically"
	// jump again from the 1-cell island; we would still fall.
	//
	// Right: the ideal scenario, *waiting 2 steps* is enough to make the
	// double jump and survive.
	// As the landing place for the second jump would be H
	// (were we to wait 2 steps), we need to make sure it is ground when
	// jumping prematurely because of !C; otherwise we should wait (i.e. not jump).
	//
	// FAIL (premature)      FAIL (imaginary)      SUCCESS (later)
	//
	// ....@............     ....@...@........     ......@...@...@..
	// .../.\...........     .../.\./.\.......     ...../.\./.\./.\.
	// ..@...@..........     ..@...@...@......     ....@...@...@...@
	// #####.#@##..#.###     #####.#.##@.#.###     #####.#.##..#.###
	// 1  ABCDEFGHI             ABCDEFGHI               ABCDEFGHI
	// 2      ABCDEFGHI             ABCDEFGHI               ABCDEFGHI
	// 3                                                        ABCD
	//
	// I used WolframAlpha to get me a different form, just for fun.
	// AND minimal form from:
	// https://www.wolframalpha.com/input/?i=%28not+A+and+D%29+or+%28not+B+and+D%29+or+%28not+C+and+D+and+H%29
	//
	// We could certainly write this with less instructions.
	//
	// J = (!A ^ D) v (!B ^ D) v (!C ^ D ^ H)
	//   = !(A ^ B ^ C) ^ !(A ^ B ^ !H) ^ D
	instructions := []rune(strings.Join([]string{
		"NOT A T", // T = !A
		"NOT T T", // T = A
		"AND B T", // T = A ^ B
		"AND C T", // T = A ^ B ^ C
		"NOT T J", // J = !(A ^ B ^ C)

		"NOT H T", // T = !H
		"AND A T", // T = A ^ !H
		"AND B T", // T = A ^ B ^ !H
		"NOT T T", // T = !(A ^ B ^ !H)

		"AND T J", // J = !(A ^ B ^ C) ^ !(A ^ B ^ !H)

		"AND D J", // J = !(A ^ B ^ C) ^ !(A ^ B ^ !H) ^ D

		"RUN",
		"\n",
	}, "\n"))
	instructionsI64 := make([]int64, len(instructions))
	for i := range instructions {
		instructionsI64[i] = int64(instructions[i])
	}

	outputI64, _ := cpu.Exec(instructionsI64)

	for i, ch := range outputI64 {
		if ch >= 0 && ch <= 255 {
			if debugPlayMovie {
				if ch == '\n' && i > 0 && outputI64[i-1] == '\n' {
					fmt.Print("\n")
					time.Sleep(500 * time.Millisecond)
					fmt.Print("\033[2J")
					continue
				}
			}

			fmt.Print(string(rune(ch)))
			continue
		}

		return int(ch)
	}

	return -1
}
