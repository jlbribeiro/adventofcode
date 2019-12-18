package main

import (
	"fmt"
	"os"

	"github.com/jlbribeiro/adventofcode/2019/day16/part2/fft"
)

func main() {
	nPhases := 100
	period := 10000
	out := fft.RealFFTFromInput(os.Stdin, nPhases, period)
	fmt.Println(out[:8])
}
