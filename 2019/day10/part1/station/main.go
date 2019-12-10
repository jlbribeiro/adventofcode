package station

import (
	"bufio"
	"io"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func gcd(a, b int) int {
	for b != 0 {
		tmp := b
		b = a % b
		a = tmp
	}

	return a
}

func AsteroidsGridFromInput(in io.Reader) [][]bool {
	var grid [][]bool

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		row := scanner.Text()

		gridRow := make([]bool, len(row))
		for i, c := range row {
			if c == '#' {
				gridRow[i] = true
			}
		}

		grid = append(grid, gridRow)
	}

	return grid
}

func MarkOutOfLoS(grid [][]bool, rx int, ry int, tx int, ty int) {
	dx, dy := tx-rx, ty-ry

	switch {
	// Same row
	case dy == 0:
		dx /= abs(dx)

	// Same column
	case dx == 0:
		dy /= abs(dy)

	// Same diagonal
	case abs(dx) == abs(dy):
		dx /= abs(dx)
		dy /= abs(dy)

	default:
		gcd := gcd(abs(dx), abs(dy))
		dx /= gcd
		dy /= gcd
	}

	x, y := tx+dx, ty+dy
	for y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y]) {
		grid[y][x] = false

		y += dy
		x += dx
	}
}

func BFSLoSAt(grid [][]bool, rx int, ry int) int {
	count := 0

	for radius := 1; ; radius++ {
		y0, x0 := ry-radius, rx-radius
		yMax, xMax := ry+radius, rx+radius

		// All coords will be outside the grid.
		if y0 < 0 && x0 < 0 && yMax >= len(grid) && xMax >= len(grid[0]) {
			break
		}

		for y := y0; y <= yMax; y++ {
			if y < 0 || y >= len(grid) {
				continue
			}

			for x := x0; x <= xMax; x++ {
				if x < 0 || x >= len(grid[y]) {
					continue
				}

				if y != y0 && y != yMax && x != x0 && x != xMax {
					continue
				}

				if !grid[y][x] {
					continue
				}

				// Asteroid in LoS
				count++
				MarkOutOfLoS(grid, rx, ry, x, y)
			}
		}
	}

	return count
}

func LoSAt(grid [][]bool, x int, y int) int {
	wgrid := make([][]bool, len(grid))
	for i := range grid {
		wgrid[i] = make([]bool, len(grid[i]))
		copy(wgrid[i], grid[i])
	}

	return BFSLoSAt(wgrid, x, y)
}

func BestLoSLocation(grid [][]bool) (int, int, int) {
	bestLoS := 0
	bestX, bestY := 0, 0

	for y := range grid {
		for x := range grid[y] {
			if !grid[y][x] {
				continue
			}

			los := LoSAt(grid, x, y)
			if los > bestLoS {
				bestLoS = los
				bestX, bestY = x, y
			}
		}
	}

	return bestLoS, bestX, bestY
}
