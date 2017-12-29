package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jlbribeiro/adventofcode/2017/day10/part1/knothash"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lengthsLine := scanner.Text()
		lengthsStr := strings.Split(lengthsLine, ",")
		lengths := make([]int, len(lengthsStr))

		for i, lengthStr := range lengthsStr {
			length64, err := strconv.ParseInt(lengthStr, 10, 64)
			if err != nil {
				panic(err)
			}
			length := int(length64)

			lengths[i] = length
		}

		hash := knothash.Hash(256, lengths)
		fmt.Println(hash)
	}
}
