package fuel

import (
	"math"
)

const GridSide int = 300

func CellPowerLevel(serialNumber int, x, y int) int {
	rackID := x + 10
	powerLevel := rackID * y
	powerLevel += serialNumber
	powerLevel *= rackID
	powerLevel = (powerLevel / 100) % 10
	return powerLevel - 5
}

func MaxTotalPower(serialNumber int) (int, int, int, int) {
	maxPowerLevel := math.MinInt32
	maxPowerX := 0
	maxPowerY := 0
	maxPowerSquareSide := 1

	// The full fuel grid is GridSide x GridSide, calculate every cell's
	// power level (the base case).
	squareSide := 1
	grids := make([][][]int, GridSide, GridSide)
	grids[squareSide-1] = make([][]int, GridSide, GridSide)
	for i := range grids[squareSide-1] {
		grids[squareSide-1][i] = make([]int, GridSide, GridSide)
		for j := range grids[squareSide-1][i] {
			grids[squareSide-1][i][j] = CellPowerLevel(serialNumber, j+1, i+1)
			if grids[squareSide-1][i][j] > maxPowerLevel {
				maxPowerLevel = grids[squareSide-1][i][j]
				maxPowerX, maxPowerY = j+1, i+1
				// maxPowerSquareSide is 1 by this point
			}
		}
	}

	for squareSide = 2; squareSide <= GridSide; squareSide++ {
		nSquaresPerSide := GridSide - squareSide + 1
		grids[squareSide-1] = make([][]int, nSquaresPerSide, nSquaresPerSide)

		for i := range grids[squareSide-1] {
			grids[squareSide-1][i] = make([]int, nSquaresPerSide, nSquaresPerSide)

			if squareSide%2 == 0 {
				// If the square side is even, that square's total power is the
				// result of combining the total power of the (previously calculated)
				// 4 squares of half the side.
				//
				// E.g.: the total fuel power in a
				// square of side 4, with a top-left corner in 3,7 is
				// the result of summing
				//   square of side 2, top-left corner in 3,7
				//   square of side 2, top-left corner in 3,9
				//   square of side 2, top-left corner in 5,7
				//   square of side 2, top-left corner in 5,9
				halfSquare := squareSide / 2
				for j := range grids[squareSide-1][i] {
					grids[squareSide-1][i][j] = grids[halfSquare-1][i][j]
					grids[squareSide-1][i][j] += grids[halfSquare-1][i][j+halfSquare]
					grids[squareSide-1][i][j] += grids[halfSquare-1][i+halfSquare][j]
					grids[squareSide-1][i][j] += grids[halfSquare-1][i+halfSquare][j+halfSquare]

					if grids[squareSide-1][i][j] > maxPowerLevel {
						maxPowerLevel = grids[squareSide-1][i][j]
						maxPowerX, maxPowerY = j+1, i+1
						maxPowerSquareSide = squareSide
					}
				}

			} else {
				// If the square side is odd, that square's total power can be
				// calculated using the biggest even-sided square covered by it,
				// summing the other cells not covered by that smaller square
				// (the last row + the last column of that square, without
				// accounting for the bottom-right corner's power level twice).
				for j := range grids[squareSide-1][i] {
					// Use the even side square total power as base.
					grids[squareSide-1][i][j] = grids[squareSide-2][i][j]

					// Sum the missing bottom row and right column.
					for k := 0; k < squareSide; k++ {
						grids[squareSide-1][i][j] += grids[0][i+squareSide-1][j+k]
						grids[squareSide-1][i][j] += grids[0][i+k][j+squareSide-1]
					}

					// Remove the bottom-right corner, as it got summed twice by
					// the loop above.
					grids[squareSide-1][i][j] -= grids[0][i+squareSide-1][j+squareSide-1]

					if grids[squareSide-1][i][j] > maxPowerLevel {
						maxPowerLevel = grids[squareSide-1][i][j]
						maxPowerX, maxPowerY = j+1, i+1
						maxPowerSquareSide = squareSide
					}
				}
			}
		}
	}

	return maxPowerLevel, maxPowerX, maxPowerY, maxPowerSquareSide
}
