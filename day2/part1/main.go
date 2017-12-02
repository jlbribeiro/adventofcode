package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	hash := int64(0)
	for scanner.Scan() {
		line := scanner.Text()

		cols := strings.Split(line, "\t")
		min := int64(math.MaxInt64)
		max := int64(math.MinInt64)

		for _, ns := range cols {
			n, err := strconv.ParseInt(ns, 10, 64)
			if err != nil {
				panic(err)
			}

			if n < min {
				min = n
			}

			if n > max {
				max = n
			}
		}

		hash += max - min
	}

	fmt.Println(hash)
}
