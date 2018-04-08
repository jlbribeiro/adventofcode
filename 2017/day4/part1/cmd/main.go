package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2017/day4/part1/pass"
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
