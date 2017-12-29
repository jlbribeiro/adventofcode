package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2017/day20/part2/particles"
)

func main() {
	lines := []string(nil)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	analyser := particles.NewAnalyserFromInput(lines)
	fmt.Println(analyser.RemainingParticles())
}
