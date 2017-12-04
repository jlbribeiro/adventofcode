package main

import (
	"bufio"
	"fmt"
	"os"

	"git.jlbribeiro.com/adventofcode/day4/part1/pass"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	nValidPasswords := 0
	for scanner.Scan() {
		passphrase := scanner.Text()
		if pass.ValidPassphrase(passphrase) {
			nValidPasswords++
		}
	}

	fmt.Println(nValidPasswords)
}
