package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/day7/part2/circus"
)

func main() {
	tower := circus.NewTower()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		programYell := scanner.Text()
		program := circus.NewProgramFromYell(programYell)
		tower.RegisterProgram(program)
	}

	fmt.Println(tower.FindWrongWeightProgramIdealWeight())
}
