package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jlbribeiro/adventofcode/2018/day5/part2/polymers"
)

func main() {
	inp, _ := ioutil.ReadAll(os.Stdin)
	input := strings.TrimSpace(string(inp))
	result := polymers.RogueUnitReact(input)
	fmt.Printf("'%s'\n", result)
	fmt.Println(len(result))
}
