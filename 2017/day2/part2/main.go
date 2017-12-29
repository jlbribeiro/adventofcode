package main

import (
	"bufio"
	"fmt"
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

		row := make([]int64, len(cols), len(cols))
		for i, ns := range cols {
			n, err := strconv.ParseInt(ns, 10, 64)
			if err != nil {
				panic(err)
			}

			row[i] = n
		}

	loop:
		for i := 0; i < len(row); i++ {
			for j := i + 1; j < len(row); j++ {
				a := row[i]
				b := row[j]
				if a > b {
					a, b = b, a
				}

				if b%a == 0 {
					hash += b / a
					break loop
				}
			}
		}
	}

	fmt.Println(hash)
}
