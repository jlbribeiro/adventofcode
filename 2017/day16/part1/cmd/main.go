package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jlbribeiro/adventofcode/day16/part1/promenade"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	steps := strings.Split(strings.TrimSpace(string(input)), ",")

	dancers := promenade.NewDancers(16)
	dancers.Dance(steps)
	fmt.Println(dancers.Alignment())
}
