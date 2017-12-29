package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2017/day18/part1/duet"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	instructions := make([]string, 0)
	for scanner.Scan() {
		instruction := scanner.Text()
		instructions = append(instructions, instruction)
	}

	musicPlayer := duet.NewMusicPlayer()
	fmt.Println(musicPlayer.Play(instructions))
}
