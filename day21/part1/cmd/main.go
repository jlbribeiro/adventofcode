package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/day21/part1/fractal"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	rules := []string(nil)
	for scanner.Scan() {
		rule := scanner.Text()
		rules = append(rules, rule)
	}

	fractalGenerator := fractal.NewGeneratorFromInput(rules)
	fractalGenerator.Run(5)
	fmt.Println(fractalGenerator.MatrixOnPixels())
}
