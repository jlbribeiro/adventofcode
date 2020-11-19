package bugs

import (
	"bufio"
	"fmt"
	"io"
)

const (
	Bug   rune = '#'
	Empty rune = '.'
)

func printBugs(state [][]bool) {
	for _, row := range state {
		for _, cell := range row {
			if cell {
				fmt.Print(string(Bug))
			} else {
				fmt.Print(string(Empty))
			}
		}
		fmt.Println()
	}
}

func RepeatedBiodiversity(firstGen [][]bool) int {
	var states [2][][]bool
	states[0] = firstGen
	states[1] = make([][]bool, len(firstGen))
	for i := range states[1] {
		states[1][i] = make([]bool, len(firstGen[i]))
	}

	cache := make(map[int]struct{})
	for i := 0; ; i++ {
		cur, next := i&1, (i+1)&1

		biodiversity := 0
		for i, row := range states[cur] {
			for j := range row {
				neighbours := 0

				for _, offsets := range [4][2]int{
					[2]int{-1, 0}, // top
					[2]int{0, -1}, // left
					[2]int{0, 1},  // right
					[2]int{1, 0},  // down
				} {
					dy, dx := offsets[0], offsets[1]
					y := i + dy
					x := j + dx
					switch {
					case y < 0,
						y >= len(states[cur]),
						x < 0,
						x >= len(states[cur][y]):
						continue
					}

					if states[cur][y][x] {
						neighbours++
					}
				}

				states[next][i][j] = states[cur][i][j]
				if states[cur][i][j] && neighbours != 1 {
					states[next][i][j] = false

				} else if !states[cur][i][j] && (neighbours == 1 || neighbours == 2) {
					states[next][i][j] = true
				}

				if states[next][i][j] {
					biodiversity |= 1 << (i*len(states[cur]) + j)
				}
			}
		}

		if _, ok := cache[biodiversity]; ok {
			return biodiversity
		}
		cache[biodiversity] = struct{}{}
	}
}

func RepeatedBiodiversityFromInput(input io.Reader) int {
	var firstGen [][]bool

	sc := bufio.NewScanner(input)
	for sc.Scan() {
		line := sc.Text()

		row := make([]bool, len(line))
		for i, ch := range line {
			row[i] = ch == Bug
		}
		firstGen = append(firstGen, row)
	}

	return RepeatedBiodiversity(firstGen)
}
