package breakout

import (
	"io"

	"github.com/jlbribeiro/adventofcode/2019/day13/part1/thermal"
)

const (
	BlockTile int = 2
)

func LevelFromInput(input io.Reader) []int64 {
	program := thermal.ProgramFromInput(input)
	cpu := thermal.NewCPU(program)
	level, _ := cpu.Exec([]int64{})
	return level
}

func BlockTilesCountFromInput(input io.Reader) int {
	level := LevelFromInput(input)
	return BlockTilesCount(level)
}

func BlockTilesCount(level []int64) int {
	tiles := len(level) / 3

	blockTiles := 0
	for i := 0; i < tiles; i++ {
		outputInstr := level[i*3 : (i+1)*3]
		if int(outputInstr[2]) == BlockTile {
			blockTiles++
		}
	}

	return blockTiles
}
