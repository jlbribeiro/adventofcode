package breakout

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/jlbribeiro/adventofcode/2019/day13/part2/thermal"
)

const (
	EmptyTile            int = 0
	WallTile             int = 1
	BlockTile            int = 2
	HorizontalPaddleTile int = 3
	BallTile             int = 4
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

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

func PrintTile(tileID int, out *strings.Builder) {
	switch tileID {
	case EmptyTile:
		out.WriteRune(' ')

	case WallTile:
		out.WriteString("\033[47m \033[0m")

	case BlockTile:
		out.WriteRune('#')

	case HorizontalPaddleTile:
		out.WriteRune('-')

	case BallTile:
		out.WriteRune('o')
	}
}

func GameScoreFromInput(programS io.Reader, printGame bool) int {
	program := thermal.ProgramFromInput(programS)
	program[0] = 2
	cpu := thermal.NewCPU(program)

	return PlayGame(cpu, printGame)
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

func PlayGame(cpu *thermal.CPU, printGame bool) int {
	score := 0

	ballX := 0
	paddleX := 0

	gameMap := make(map[[2]int]int)
	minX, minY, maxX, maxY := 0, 0, 0, 0

	var input []int64
	var output []int64
	running := true
	for running {
		output, running = cpu.Exec(input)
		nTiles := len(output) / 3

		for i := 0; i < nTiles; i++ {
			instruction := output[i*3 : (i+1)*3]
			x, y, tileID := int(instruction[0]), int(instruction[1]), int(instruction[2])

			if x == -1 && y == 0 {
				score = tileID
				continue
			}

			minX, minY, maxX, maxY = min(minX, x), min(minY, y), max(maxX, x), max(maxY, y)
			gameMap[[2]int{x, y}] = tileID

			switch tileID {
			case HorizontalPaddleTile:
				paddleX = x

			case BallTile:
				ballX = x
			}
		}

		paddleDir := ballX - paddleX
		if paddleDir != 0 {
			paddleDir /= abs(paddleDir)
		}

		input = append(input, int64(paddleDir))

		if printGame {
			gameOutput := &strings.Builder{}
			gameOutput.WriteString("\033[H\033[2J")

			for y := minY; y <= maxY; y++ {
				for x := minX; x <= maxX; x++ {
					PrintTile(gameMap[[2]int{x, y}], gameOutput)
				}
				gameOutput.WriteRune('\n')
			}
			gameOutput.WriteRune('\n')
			gameOutput.WriteString(fmt.Sprintf("Score: %d\n", score))

			fmt.Print(gameOutput.String())
			time.Sleep(50 * time.Millisecond)
		}
	}

	return score
}
