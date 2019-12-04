package wirejumble

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

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

type Direction rune

const (
	Up    Direction = 'U'
	Down            = 'D'
	Left            = 'L'
	Right           = 'R'
)

func (d Direction) Deltas() (int, int) {
	switch d {
	case Up:
		return 0, -1
	case Down:
		return 0, 1
	case Left:
		return -1, 0
	case Right:
		return 1, 0
	default:
		panic(fmt.Errorf("Invalid direction"))
	}
}

type WireSegment struct {
	Length    int
	Direction Direction
}

func (s WireSegment) String() string {
	return fmt.Sprintf("%c%d", s.Direction, s.Length)
}

type Wire struct {
	Segments []WireSegment
}

type Wires []*Wire

func IntersectionDistanceFromInput(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	var wires Wires
	for scanner.Scan() {
		inputLine := scanner.Text()
		segsS := strings.Split(inputLine, ",")

		wire := &Wire{}
		for _, segS := range segsS {
			direction, lengthS := segS[:1], segS[1:]
			length64, err := strconv.ParseInt(lengthS, 10, 32)
			if err != nil {
				panic(err)
			}

			segment := WireSegment{
				Length:    int(length64),
				Direction: Direction(direction[0]),
			}

			wire.Segments = append(wire.Segments, segment)
		}

		wires = append(wires, wire)
	}

	return wires.IntersectionDistance()
}

func (wires *Wires) Boundaries() (int, int, int, int) {
	minX, maxX := 0, 0
	minY, maxY := 0, 0

	for _, wire := range *wires {
		x, y := 0, 0
		for _, seg := range wire.Segments {
			dX, dY := seg.Direction.Deltas()
			x += seg.Length * dX
			y += seg.Length * dY

			minX = min(minX, x)
			maxX = max(maxX, x)
			minY = min(minY, y)
			maxY = max(maxY, y)
		}
	}

	return minX, maxX, minY, maxY
}

func (wires *Wires) IntersectionDistance() int {
	minX, maxX, minY, maxY := wires.Boundaries()

	width := maxX - minX + 1
	height := maxY - minY + 1
	grid := make([][]int, height, height)
	for i := range grid {
		grid[i] = make([]int, width, width)
	}

	// "Origin" coordinates offsets.
	// Translate "real" coordinates into grid coordinates.
	origX, origY := -minX, -minY

	minIntersectDist := math.MaxInt32
	for i, wire := range *wires {
		wireID := 1 << i
		x, y := 0, 0
		for _, seg := range wire.Segments {
			dX, dY := seg.Direction.Deltas()
			for j := 0; j < seg.Length; j++ {
				x, y = x+dX, y+dY

				// First wire occupying this cell
				if grid[origY+y][origX+x] == 0 {
					grid[origY+y][origX+x] = wireID
					continue
				}

				// Wire intersecting itself
				if grid[origY+y][origX+x]&wireID == wireID {
					continue
				}

				// Second wire intersecting first wire
				intersectDist := abs(y) + abs(x)
				minIntersectDist = min(minIntersectDist, intersectDist)
			}
		}
	}

	return minIntersectDist
}
