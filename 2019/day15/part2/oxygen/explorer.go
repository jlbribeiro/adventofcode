package oxygen

import (
	"fmt"
	"io"
	"math"

	"github.com/jlbribeiro/adventofcode/2019/day15/part2/thermal"
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

type Status int

const (
	HitWall     Status = 0
	Moved       Status = 1
	FoundOxygen Status = 2
)

type Block int

const (
	Wall   Block = -1
	Space  Block = -2
	Oxygen Block = -3
)

func StatusToBlock(status Status) Block {
	switch status {
	case HitWall:
		return Wall

	case Moved:
		return Space

	case FoundOxygen:
		return Oxygen

	default:
		panic(fmt.Errorf("unexpected Status: %d", int(status)))
	}
}

type Cell struct {
	Block            Block
	DistanceToOrigin int
}

type Explorer struct {
	x, y             int
	oxygenX, oxygenY int

	walker Walker
	area   map[[2]int]Cell
}

func NewExplorer(walker Walker) *Explorer {
	area := make(map[[2]int]Cell)
	area[[2]int{0, 0}] = Cell{
		Block:            Space,
		DistanceToOrigin: 0,
	}

	return &Explorer{
		x:      0,
		y:      0,
		walker: walker,
		area:   area,
	}
}

func (exp *Explorer) DistanceToOrigin(x, y int) int {
	cell, ok := exp.area[[2]int{x, y}]
	if !ok {
		return math.MaxInt64
	}

	return cell.DistanceToOrigin
}

func (exp *Explorer) Walk(direction Direction, distance int) Status {
	status := exp.walker.Walk(direction)

	dx, dy := direction.Offsets()
	x, y := exp.x+dx, exp.y+dy

	cellDist := min(exp.DistanceToOrigin(x, y), distance+1)
	exp.area[[2]int{x, y}] = Cell{
		Block:            StatusToBlock(status),
		DistanceToOrigin: cellDist,
	}

	if status == HitWall {
		return status
	}

	if status == FoundOxygen {
		exp.oxygenX, exp.oxygenY = x, y
	}

	exp.x, exp.y = x, y
	return status
}

func (exp *Explorer) ShouldExplore(direction Direction, distance int) bool {
	dx, dy := direction.Offsets()
	x, y := exp.x+dx, exp.y+dy
	cell, explored := exp.area[[2]int{x, y}]

	if !explored {
		return true
	}

	// NOTE: For my input, this condition is always false; that's a mix of luck
	// with the fact that the maze is essentially a "narrow" path.
	// Wanting to make sure this edge condition did indeed exist lead to
	// implementing Walker (so I could mock it), to feed it an "open" maze
	// (where the same cell can be reached via a multitude of paths.
	return cell.DistanceToOrigin > distance+1
}

func (exp *Explorer) NextDirections(distance int) []Direction {
	directions := []Direction{}

	for direction := North; direction <= East; direction++ {
		if exp.ShouldExplore(direction, distance) {
			directions = append(directions, direction)
		}
	}

	return directions
}

func (exp *Explorer) Explore(distance int) {
	directions := exp.NextDirections(distance)
	for _, direction := range directions {
		status := exp.Walk(direction, distance)
		if status == HitWall {
			continue
		}

		exp.Explore(distance + 1)

		// distance argument doesn't matter
		exp.Walk(direction.Opposite(), distance+1)
	}
}

func (exp *Explorer) ExploreMaze() {
	exp.Explore(0)
}

func (exp *Explorer) ShortestPathToOxygen() int {
	return exp.DistanceToOrigin(exp.oxygenX, exp.oxygenY)
}

func (exp *Explorer) PrintMaze() {
	minX, minY, maxX, maxY := 0, 0, 0, 0
	for coords := range exp.area {
		x, y := coords[0], coords[1]
		minX = min(minX, x)
		minY = min(minY, y)
		maxX = max(maxX, x)
		maxY = max(maxY, y)
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			cell, ok := exp.area[[2]int{x, y}]
			if !ok {
				fmt.Printf("\033[47m \033[0m")
				continue
			}

			if x == 0 && y == 0 {
				fmt.Printf("\033[33mX\033[0m")
				continue
			}

			switch cell.Block {
			case Wall:
				fmt.Printf("\033[47m \033[0m")

			case Space:
				// fmt.Printf(" ")
				fmt.Printf("%d", cell.DistanceToOrigin%10)

			case Oxygen:
				fmt.Printf("\033[31mO\033[0m")
			}
		}
		fmt.Println()
	}
}

func FindShortestPathToOxygenFromInput(in io.Reader, printMaze bool) int {
	program := thermal.ProgramFromInput(in)
	cpu := thermal.NewCPU(program)
	cpuWalker := NewCPUWalker(cpu)
	explorer := NewExplorer(cpuWalker)

	explorer.ExploreMaze()

	if printMaze {
		explorer.PrintMaze()
		fmt.Println()
	}

	return explorer.ShortestPathToOxygen()

}

func (exp *Explorer) OxygenPropagation() int {
	// Reuse the maze (cell.Block info) from part1, just discard distance info.
	for coords, cell := range exp.area {
		cell.DistanceToOrigin = -1
		exp.area[coords] = cell
	}

	queue := make([][2]int, 0, len(exp.area))
	queue = append(queue, [2]int{exp.oxygenX, exp.oxygenY})
	distance := 0

	for ; len(queue) > 0; distance++ {
		nCells := len(queue)
		for i := 0; i < nCells; i++ {
			coords := queue[0]

			if exp.area[coords].DistanceToOrigin >= 0 {
				queue = queue[1:]
				continue
			}

			x, y := coords[0], coords[1]

			cell := exp.area[coords]
			cell.DistanceToOrigin = distance
			exp.area[coords] = cell

			for direction := North; direction <= East; direction++ {
				dx, dy := direction.Offsets()
				rx, ry := x+dx, y+dy
				coords := [2]int{rx, ry}
				if exp.area[coords].Block == Wall {
					continue
				}

				if exp.area[coords].DistanceToOrigin < 0 {
					queue = append(queue, [2]int{rx, ry})
				}
			}

			queue = queue[1:]
		}
	}

	distance--
	return distance
}

func OxygenPropagationFromInput(in io.Reader, printMaze bool) int {
	program := thermal.ProgramFromInput(in)
	cpu := thermal.NewCPU(program)
	cpuWalker := NewCPUWalker(cpu)
	explorer := NewExplorer(cpuWalker)
	explorer.ExploreMaze()

	oxygenPropagation := explorer.OxygenPropagation()

	if printMaze {
		explorer.PrintMaze()
		fmt.Println()
	}

	return oxygenPropagation

}
