package asciicamera

import (
	"io"

	"github.com/jlbribeiro/adventofcode/2019/day11/part1/thermal"
)

const scaffold rune = '#'

func SumAlignmentCoordinates(photo [][]rune) int {
	for i := range photo {
		for j := range photo[i] {
			switch photo[i][j] {
			case '^', 'v', '<', '>':
				photo[i][j] = scaffold
			}
		}
	}

	parameters := 0
	for y := 1; y < len(photo)-1; y++ {
		for x := 1; x < len(photo[y])-1; x++ {
			intersect := photo[y][x] == scaffold
			intersect = intersect && photo[y-1][x] == scaffold
			intersect = intersect && photo[y+1][x] == scaffold
			intersect = intersect && photo[y][x-1] == scaffold
			intersect = intersect && photo[y][x+1] == scaffold

			if intersect {
				parameters += y * x
			}
		}
	}

	return parameters
}

func SumAlignmentCoordinatesFromInput(input io.Reader) int {
	program := thermal.ProgramFromInput(input)
	cpu := thermal.NewCPU(program)

	photoA, _ := cpu.Exec([]int64{})

	var photo [][]rune
	var row []rune
	for _, pixel := range photoA {
		if rune(pixel) == '\n' && len(row) > 0 {
			photo = append(photo, row)
			row = nil
			continue
		}

		row = append(row, rune(pixel))
	}

	return SumAlignmentCoordinates(photo)
}
