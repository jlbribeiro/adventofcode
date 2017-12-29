package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jlbribeiro/adventofcode/day18/part2/duet"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	instructions := make([]string, 0)
	for scanner.Scan() {
		instruction := strings.TrimSpace(scanner.Text())
		instructions = append(instructions, instruction)
	}

	duet := duet.NewDuet()
	fmt.Println(duet.Play(instructions))
}
