package cryo

import (
	"bufio"
	"fmt"
	"io"

	"github.com/jlbribeiro/adventofcode/2019/day25/part1/intcode"
)

func WalkFromInput(programInput io.Reader, commandsInput io.Reader) {
	program := intcode.ProgramFromInput(programInput)
	cpu := intcode.NewCPU(program)

	reader := bufio.NewReader(commandsInput)
	var inputBuffer []int64
	for {
		output, waitingInput := cpu.Exec(inputBuffer)
		for _, ch := range output {
			fmt.Print(string(ch))
		}

		if !waitingInput {
			return
		}

		fmt.Print("> ")
		command, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}

		inputBuffer = []int64{}
		for _, ch := range command {
			inputBuffer = append(inputBuffer, int64(ch))
		}
		inputBuffer = append(inputBuffer, int64('\n'))
	}
}
