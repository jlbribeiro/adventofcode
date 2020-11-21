package bugs

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

const (
	Size   int = 5
	Center int = Size / 2
)

const (
	Bug   rune = '#'
	Empty rune = '.'
)

func printBugs(state *[Size][Size]bool) {
	for y, row := range state {
		for x, cell := range row {
			if y == Center && x == Center {
				fmt.Print(string("?"))
				continue
			}

			if cell {
				fmt.Print(string(Bug))
			} else {
				fmt.Print(string(Empty))
			}
		}
		fmt.Println()
	}
}

type Bugs struct {
	nGenerations int
	states       [2]map[int]*[Size][Size]bool
	minLevel     int
	maxLevel     int
}

func NewBugs(initialState [Size][Size]bool) *Bugs {
	var states [2]map[int]*[Size][Size]bool

	for i := range states {
		states[i] = make(map[int]*[Size][Size]bool)
	}

	nGenerations := 0
	curState := nGenerations & 1
	nextState := (nGenerations + 1) & 1
	level := 0

	states[curState][level] = &initialState
	states[nextState][level] = &[Size][Size]bool{}

	return &Bugs{
		nGenerations: nGenerations,
		states:       states,
		minLevel:     0,
		maxLevel:     0,
	}
}

func (b *Bugs) get(level int, y int, x int) bool {
	state, ok := b.states[b.nGenerations&1][level]
	if !ok {
		return false
	}

	return state[y][x]
}

func (b *Bugs) countRow(level int, y int) int {
	count := 0
	for x := 0; x < Size; x++ {
		if b.get(level, y, x) {
			count++
		}
	}
	return count
}

func (b *Bugs) countCol(level int, x int) int {
	count := 0
	for y := 0; y < Size; y++ {
		if b.get(level, y, x) {
			count++
		}
	}
	return count
}

func (b *Bugs) step() {
	debug := false

	curGen := b.nGenerations & 1
	nextGen := (b.nGenerations + 1) & 1

	if debug {
		fmt.Printf("Generation %d\n", b.nGenerations+1)
	}

	minLevel, maxLevel := b.minLevel-1, b.maxLevel+1
	for level := minLevel; level <= maxLevel; level++ {
		var curState *[Size][Size]bool
		var nextState *[Size][Size]bool

		isBoundaryLevel := level < b.minLevel || level > b.maxLevel
		if isBoundaryLevel {
			curState = &[Size][Size]bool{}
			nextState = &[Size][Size]bool{}

		} else {
			curState = b.states[curGen][level]
			nextState = b.states[nextGen][level]
		}

		if debug {
			fmt.Printf("level: %d\n", level)
			fmt.Printf("(contains level %d)\n", level+1)
			fmt.Println("before")
			printBugs(curState)
			fmt.Println()
		}

		nBugs := 0
		for y := 0; y < Size; y++ {
			for x := 0; x < Size; x++ {
				if y == Center && x == Center {
					if debug {
						fmt.Print("?")
					}
					continue
				}

				nNeighbours := 0
				for _, offsets := range [4][2]int{
					[2]int{-1, 0}, // top
					[2]int{0, -1}, // left
					[2]int{0, 1},  // right
					[2]int{1, 0},  // down
				} {
					dy, dx := offsets[0], offsets[1]
					ny := y + dy
					nx := x + dx
					switch {
					// Level below
					case nx == Center && ny == Center:
						// One of the cell (y,x) neighbours is the center cell;
						// the exact cells of the level below that should be
						// counted depend on the (y, x) coordinates.
						switch {
						case y == Center-1 && x == Center:
							nNeighbours += b.countRow(level+1, 0)

						case y == Center && x == Center-1:
							nNeighbours += b.countCol(level+1, 0)

						case y == Center && x == Center+1:
							nNeighbours += b.countCol(level+1, Size-1)

						case y == Center+1 && x == Center:
							nNeighbours += b.countRow(level+1, Size-1)

						default:
							panic("unexpected condition")
						}

					// Level above
					case ny < 0:
						// Level above, cell 8 (1, 2)
						if b.get(level-1, Center-1, Center) {
							nNeighbours++
						}

					case nx < 0:
						// Level above, cell 12 (2, 1)
						if b.get(level-1, Center, Center-1) {
							nNeighbours++
						}

					case nx >= Size:
						// Level above, cell 14 (2, 3)
						if b.get(level-1, Center, Center+1) {
							nNeighbours++
						}

					case ny >= Size:
						// Level above, cell 18 (3, 2)
						if b.get(level-1, Center+1, Center) {
							nNeighbours++
						}

					default:
						if curState[ny][nx] {
							nNeighbours++
						}
					}
				}

				nextState[y][x] = curState[y][x]
				if curState[y][x] && nNeighbours != 1 {
					nextState[y][x] = false

				} else if !curState[y][x] && (nNeighbours == 1 || nNeighbours == 2) {
					nextState[y][x] = true
				}

				if nextState[y][x] {
					nBugs++
				}

				if debug {
					fmt.Print(string(strconv.FormatInt(int64(nNeighbours), 10)))
				}
			}

			if debug {
				fmt.Println()
			}
		}

		if debug {
			fmt.Println()
			fmt.Println("after")
			printBugs(nextState)
			fmt.Printf("nBugs: %d\n", nBugs)
			fmt.Println()
		}

		if isBoundaryLevel && nBugs > 0 {
			// TODO: should create new states and update min/max level
			b.states[nextGen][level] = nextState
			b.states[curGen][level] = curState
			if level < b.minLevel {
				b.minLevel = level

			} else if level > b.maxLevel {
				b.maxLevel = level

			} else {
				panic("unexpected condition: isBoundaryLevel but not outside min/max levels")
			}
		}
	}

	b.nGenerations++

	if debug {
		fmt.Println("###################################")
		fmt.Println()
	}

}

func (b *Bugs) stepN(nGenerations int) {
	for i := 0; i < nGenerations; i++ {
		b.step()
	}
}

func (b *Bugs) countBugs() int {
	count := 0
	for level := b.minLevel; level <= b.maxLevel; level++ {
		for _, row := range b.states[b.nGenerations&1][level] {
			for _, cell := range row {
				if cell {
					count++
				}
			}
		}
	}

	return count
}

func BugsAfterNGenerations(firstGen [][]bool, nGenerations int) int {
	var states [2][][]bool
	states[0] = firstGen
	states[1] = make([][]bool, len(firstGen))
	for i := range states[1] {
		states[1][i] = make([]bool, len(firstGen[i]))
	}

	for i := 0; i < nGenerations; i++ {
		cur, next := i&1, (i+1)&1

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
					// Level below
					case x == Center && y == Center:
						// TODO: this is the case when one of
						// cell (i,j) neighbours is the center cell;
						// the exact cells of the level below that should be
						// counted depend on the (i, j) coordinates.
						continue

					// Level above
					case y < 0:
						// TODO: level above, cell 8 (1, 2)
						continue

					case x < 0:
						// TODO: level above, cell 12 (2, 1)
						continue

					case x >= Size:
						// TODO: level above, cell 14 (2, 3)
						continue

					case y >= Size:
						// TODO: level above, cell 18 (3, 2)
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
			}
		}
	}

	count := 0
	lastGen := (nGenerations + 1) & 1
	for _, row := range states[lastGen] {
		for _, cell := range row {
			if cell {
				count++
			}
		}
	}
	return count
}

func BugsAfterNGenerationsFromInput(input io.Reader, nGenerations int) int {
	var firstGen [5][5]bool

	sc := bufio.NewScanner(input)
	row := 0
	for sc.Scan() {
		line := sc.Text()

		for col, ch := range line {
			firstGen[row][col] = ch == Bug
		}
		row++
	}

	bugs := NewBugs(firstGen)
	bugs.stepN(nGenerations)
	return bugs.countBugs()
}
