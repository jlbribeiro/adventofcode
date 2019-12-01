package fuel

import (
	"bufio"
	"io"
	"strconv"
)

var fuelCache map[int]int

func init() {
	fuelCache = make(map[int]int)
}

func ModuleRequirements(mass int) int {
	return mass/3 - 2
}

func TotalRequirementsFromInput(in io.Reader) int {
	scanner := bufio.NewScanner(in)
	total := 0
	for scanner.Scan() {
		massS := scanner.Text()
		mass64, err := strconv.ParseInt(massS, 10, 32)
		if err != nil {
			panic(err)
		}

		total += ModuleRequirements(int(mass64))
	}

	return total
}
