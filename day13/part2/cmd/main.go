package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jlbribeiro/adventofcode/day13/part2/firewall"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fw := firewall.NewFirewall()
	for scanner.Scan() {
		line := scanner.Text()
		depthRange := strings.Split(line, ": ")

		depth, err := strconv.ParseInt(depthRange[0], 10, 64)
		if err != nil {
			panic(err)
		}

		scanningRange, err := strconv.ParseInt(depthRange[1], 10, 64)
		if err != nil {
			panic(err)
		}

		fw.AddLayer(firewall.NewLayer(int(depth), int(scanningRange)))
	}

	fmt.Println(fw.UndetectedWalkDelay())
}
