package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		l := len(line)

		sum := int64(0)
		for i := range line {
			if line[i] == line[(i+1)%l] {
				d, err := strconv.ParseInt(string(line[i]), 10, 64)
				if err != nil {
					panic(err)
				}

				sum += d
			}
		}

		fmt.Println(sum)
	}
}
