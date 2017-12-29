package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2017/day21/part2/fractal"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	rules := []string(nil)
	for scanner.Scan() {
		rule := scanner.Text()
		rules = append(rules, rule)
	}

	fractalGenerator := fractal.NewGeneratorFromInput(rules)
	fractalGenerator.Run(18)
	fmt.Println(fractalGenerator.MatrixOnPixels())
}
