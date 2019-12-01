package fuel

import (
	"bufio"
	"io"
	"strconv"
)

type FuelCache map[int]int

func NewFuelCache() FuelCache {
	return make(map[int]int)
}

func (cache FuelCache) ModuleRequirements(mass int) int {
	if totalFuel, ok := cache[mass]; ok {
		return totalFuel
	}

	fuelReq := mass/3 - 2
	if fuelReq <= 0 {
		cache[mass] = 0
		return cache[mass]
	}

	cache[mass] = fuelReq + cache.ModuleRequirements(fuelReq)
	return cache[mass]
}

func TotalRequirementsFromInput(in io.Reader) int {
	scanner := bufio.NewScanner(in)

	cache := NewFuelCache()
	total := 0

	for scanner.Scan() {
		massS := scanner.Text()
		mass64, err := strconv.ParseInt(massS, 10, 32)
		if err != nil {
			panic(err)
		}

		total += cache.ModuleRequirements(int(mass64))
	}

	return total
}
