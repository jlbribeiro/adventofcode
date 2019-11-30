package fuel

import "math"

const GridSize int = 300
const SquareSize int = 3

func CellPowerLevel(serialNumber int, x, y int) int {
	rackID := x + 10
	powerLevel := rackID * y
	powerLevel += serialNumber
	powerLevel *= rackID
	powerLevel = (powerLevel / 100) % 10
	return powerLevel - 5
}

func MaxTotalPower(serialNumber int) (int, int, int) {
	maxPowerLevel := math.MinInt32
	maxPowerX := 0
	maxPowerY := 0

	for cornerY := 0; cornerY <= GridSize-SquareSize; cornerY++ {
		for cornerX := 0; cornerX <= GridSize-SquareSize; cornerX++ {
			squarePower := 0
			for y := cornerY; y < cornerY+SquareSize; y++ {
				for x := cornerX; x < cornerX+SquareSize; x++ {
					squarePower += CellPowerLevel(serialNumber, x, y)
				}
			}
			if squarePower > maxPowerLevel {
				maxPowerLevel = squarePower
				maxPowerX, maxPowerY = cornerX, cornerY
			}
		}
	}

	return maxPowerLevel, maxPowerX, maxPowerY
}
