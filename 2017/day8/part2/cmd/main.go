package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2017/day8/part2/registers"
)

func main() {
	textProgram := []string(nil)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		textInstruction := scanner.Text()
		textProgram = append(textProgram, textInstruction)
	}

	cpu := registers.NewCPUFromTextInstructions(textProgram)
	cpu.RunProgram()
	fmt.Println(cpu.GetLargestRegisterValueEver())
}
