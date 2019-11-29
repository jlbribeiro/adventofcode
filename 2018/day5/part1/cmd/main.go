package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jlbribeiro/adventofcode/2018/day5/part1/polymers"
)

func main() {
	inp, _ := ioutil.ReadAll(os.Stdin)
	input := strings.TrimSpace(string(inp))
	result := polymers.React(input)
	fmt.Printf("'%s'\n", result)
	fmt.Println(len(result))
}
