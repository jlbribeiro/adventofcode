package tubes

import (
	"fmt"
	"strings"
)

type Walker struct {
	grid      [][]rune
	pos       *Vector
	dir       Direction
	steps     int
	collected []rune
}

func NewWalkerFromInput(input string) *Walker {
	lines := strings.Split(input, "\n")

	grid := make([][]rune, len(lines))
	for i := range grid {
		grid[i] = []rune(lines[i])
	}

	y := 0
	x := 0
	for ; x < len(grid[y]); x++ {
		if grid[y][x] == '|' {
			break
		}
	}

	pos := &Vector{y, x}

	return &Walker{
		grid:      grid,
		pos:       pos,
		dir:       SOUTH,
		steps:     1,
		collected: make([]rune, 0),
	}
}

func (w *Walker) at(pos *Vector) rune {
	if pos.y < 0 || pos.y >= len(w.grid) || pos.x < 0 || pos.x >= len(w.grid[pos.y]) {
		return ' '
	}

	return w.grid[pos.y][pos.x]
}

func (w *Walker) step(dir Direction) *Vector {
	return w.pos.Add(dir.Vector())
}

func expectedTrack(dir Direction) rune {
	if dir == NORTH || dir == SOUTH {
		return '|'
	} else if dir == WEST || dir == EAST {
		return '-'
	}

	return ' '
}

func (w *Walker) chooseNextDirection() {
	left := w.dir.RotateLeft()
	right := w.dir.RotateRight()

	lookLeft := w.at(w.step(left))
	lookRight := w.at(w.step(right))

	expectedTrack := expectedTrack(left)

	if lookLeft == expectedTrack || lookLeft != ' ' {
		w.dir = left
	} else if lookRight == expectedTrack || lookRight != ' ' {
		w.dir = right
	} else {
		fmt.Println(fmt.Sprintf("Currently in %v", string(w.at(w.pos))))
	}
}

func (w *Walker) Walk() string {
	for {
		// Walk following current direction
		w.pos = w.step(w.dir)

		ch := w.at(w.pos)

		if ch == '|' || ch == '-' {
			w.steps++
			continue
		}

		if ch == '+' {
			w.steps++
			w.chooseNextDirection()
			continue
		}

		if ch == ' ' {
			break
		}

		// Collecting letter
		w.steps++
		w.collected = append(w.collected, ch)
	}

	return string(w.collected)
}

func (w *Walker) Steps() int {
	return w.steps
}
