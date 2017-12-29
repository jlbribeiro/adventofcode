package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jlbribeiro/adventofcode/2017/day22/part2/virus"
)

func main() {
	infectionStatus, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	av := virus.NewAVFromInput(string(infectionStatus))
	av.Run(10000000)
	fmt.Println(av.NInfections())
}
