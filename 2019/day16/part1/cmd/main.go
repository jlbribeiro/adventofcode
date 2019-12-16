package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day16/part1/fft"
)

func main() {
	out := fft.FFTFromInput(os.Stdin, 100)
	fmt.Println(out[:8])
}
