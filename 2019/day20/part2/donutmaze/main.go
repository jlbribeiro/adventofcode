package donutmaze

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"time"
)

const labelSize int = 2

const (
	Space      rune = ' '
	Wall       rune = '#'
	Path       rune = '.'
	PortalDoor rune = '@'
)

func isVoid(cell rune) bool {
	switch cell {
	case Wall,
		Path,
		PortalDoor:
		return false
	}

	return true
}

type vec2 struct {
	y int
	x int
}

func (v vec2) String() string {
	return fmt.Sprintf("<y: %d, x: %d>", v.y, v.x)
}

type Portal struct {
	label  string
	inner  bool
	coords vec2
}

func (p Portal) String() string {
	return fmt.Sprintf("%s [inner: %s]: %s", p.label, strconv.FormatBool(p.inner), p.coords.String())
}

func findDimensions(maze [][]rune) (int, int, int) {
	height := len(maze) - 2*labelSize
	width := len(maze[0]) - 2*labelSize
	thickness := 0

	for _, row := range maze {
		nCells := 0
		for x, cell := range row {
			if cell == Wall || cell == Path {
				// Torus
				nCells++
				continue

			}

			// Labels, either inside or outside the torus.

			if nCells == 0 {
				// Top/bottom label space
				// OR
				// Left label space
				continue
			}

			// Maze row

			if nCells == width {
				// "Complete" row
				continue
			}

			// Incomplete ("holed") row;
			// since the current cell is not in torus-space, it might be
			// the inside-the-torus label space,
			// or the right label space.

			if x == labelSize+width {
				// Right label space; we can calculate thickness
				thickness = nCells / 2
			}
		}
	}

	return height, width, thickness
}

func extractLabel(maze [][]rune, coords [2]int, get func(d0 int, d1 int) rune) string {
	var label []rune
	for off := 0; off < labelSize; off++ {
		label = append(label, get(coords[0]+off, coords[1]))
	}
	return string(label)
}

func findPortals(maze [][]rune, dimSize [2]int, thickness int, get func(d0 int, d1 int) rune, realCoords func(d0 int, d1 int) vec2, portalByCoords map[vec2]Portal) {
	outerBefore := labelSize
	innerBefore := labelSize + thickness - 1
	innerAfter := dimSize[0] + labelSize - thickness
	outerAfter := labelSize + dimSize[0] - 1

	for i, d0 := range []int{outerBefore, innerBefore, innerAfter, outerAfter} {
		d1First := labelSize
		d1Last := labelSize + dimSize[1]

		innerPortals := i == 1 || i == 2
		if innerPortals {
			d1First = labelSize + thickness
			d1Last = labelSize + dimSize[1] - thickness
		}

		for d1 := d1First; d1 < d1Last; d1++ {
			if get(d0, d1) != Path {
				continue
			}

			var labelOff int

			labelBefore := i%2 == 0
			if labelBefore {
				labelOff = d0 - labelSize
			} else {
				labelOff = d0 + 1
			}

			refCoords := [2]int{labelOff, d1}
			label := extractLabel(maze, refCoords, get)

			absCoords := realCoords(d0, d1)
			portalByCoords[absCoords] = Portal{
				label:  label,
				inner:  innerPortals,
				coords: absCoords,
			}
		}
	}
}

type state struct {
	coords vec2
	nSteps int
	level  int
}

func printMazeWalk(maze [][]rune, portalByCoords map[vec2]Portal, queueSize int, state state, startPortalLabel string, endPortalLabel string) {
	var sb strings.Builder
	sb.WriteString("\033[2J")
	for r, row := range maze {
		for c, cell := range row {
			if r == state.coords.y && c == state.coords.x {
				sb.WriteString("\033[1;31mx\033[0m")
				continue
			} else if cell == PortalDoor {
				if isPortalAvailable(maze, portalByCoords, vec2{y: r, x: c}, state.level, startPortalLabel, endPortalLabel) {
					sb.WriteString(string(PortalDoor))
				} else {
					sb.WriteString(string(Wall))
				}
				continue
			}

			sb.WriteString(string(cell))
		}
		sb.WriteString("\n")
	}
	sb.WriteString(fmt.Sprintf("\nqueueSize: %d\n", queueSize))
	sb.WriteString(fmt.Sprintf("nSteps: %d\n", state.nSteps))
	sb.WriteString(fmt.Sprintf("level: %d\n", state.level))
	fmt.Println(sb.String())
	time.Sleep(150 * time.Millisecond)
}

type visit struct {
	level  int
	coords vec2
}

func isPortalAvailable(maze [][]rune, portalByCoords map[vec2]Portal, coords vec2, level int, startPortalLabel string, exitPortalLabel string) bool {
	if maze[coords.y][coords.x] != PortalDoor {
		return false
	}

	portal := portalByCoords[coords]

	if level == 0 && portal.label == startPortalLabel {
		return false
	}

	isStartExitPortal := portal.label == exitPortalLabel || portal.label == startPortalLabel
	if level > 0 && isStartExitPortal {
		return false
	}

	outerPortal := !portal.inner
	if level == 0 && outerPortal && !isStartExitPortal {
		return false
	}

	if level > 0 && isStartExitPortal {
		return false
	}

	return true
}

func bfsFromTo(maze [][]rune, portalByCoords map[vec2]Portal, portalsByLabel map[string][]Portal, startPortalLabel string, exitPortalLabel string) int {
	debugPrint := false

	visitMap := make(map[visit]struct{})
	queue := []state{state{coords: portalsByLabel[startPortalLabel][0].coords, nSteps: 0, level: 0}}
	var curState state
	for len(queue) > 0 {
		curState, queue = queue[0], queue[1:]
		if _, visited := visitMap[visit{level: curState.level, coords: curState.coords}]; visited {
			continue
		}
		visitMap[visit{level: curState.level, coords: curState.coords}] = struct{}{}

		if debugPrint {
			printMazeWalk(maze, portalByCoords, len(queue), curState, startPortalLabel, exitPortalLabel)
		}

		if maze[curState.coords.y][curState.coords.x] == PortalDoor && !(curState.level == 0 && portalByCoords[curState.coords].label == startPortalLabel) {
			portal := portalByCoords[curState.coords]

			if curState.level == 0 && portal.label == exitPortalLabel {
				return curState.nSteps
			}

			var nextCoords *vec2
			for i := range portalsByLabel[portal.label] {
				coords := portalsByLabel[portal.label][i].coords
				if coords == curState.coords {
					continue
				}

				nextCoords = &coords
			}
			if nextCoords == nil {
				panic(fmt.Sprintf("couldn't find the portal exit for %s", portal))
			}

			var nextLevel int
			if portal.inner {
				nextLevel = curState.level + 1
			} else {
				nextLevel = curState.level - 1
			}

			queue = append(queue, state{coords: *nextCoords, nSteps: curState.nSteps + 1, level: nextLevel})
		}

		for _, direction := range []Direction{North, South, West, East} {
			dy, dx := direction.Offsets()
			ny, nx := curState.coords.y+dy, curState.coords.x+dx
			if ny < 0 || ny >= len(maze) || nx < 0 || nx >= len(maze[ny]) {
				continue
			}

			if isVoid(maze[ny][nx]) || maze[ny][nx] == Wall {
				continue
			}

			if maze[ny][nx] == PortalDoor && !isPortalAvailable(maze, portalByCoords, vec2{y: ny, x: nx}, curState.level, startPortalLabel, exitPortalLabel) {
				continue
			}

			nextCoords := vec2{y: ny, x: nx}
			queue = append(queue, state{coords: nextCoords, nSteps: curState.nSteps + 1, level: curState.level})
		}
	}

	return math.MaxInt64
}

func MinStepsFromInput(input io.Reader) int {
	var maze [][]rune
	reader := bufio.NewScanner(input)
	for reader.Scan() {
		line := reader.Text()
		maze = append(maze, []rune(line))
	}

	height, width, thickness := findDimensions(maze)

	portalByCoords := make(map[vec2]Portal)

	// Find portals on both "horizontal" and "vertical" edges of the torus
	// (both outer and inner ones).
	findPortals(maze, [2]int{height, width}, thickness, func(d0 int, d1 int) rune {
		return maze[d0][d1]
	}, func(d0 int, d1 int) vec2 {
		return vec2{y: d0, x: d1}
	}, portalByCoords)
	findPortals(maze, [2]int{width, height}, thickness, func(d0 int, d1 int) rune {
		return maze[d1][d0]
	}, func(d0 int, d1 int) vec2 {
		return vec2{y: d1, x: d0}
	}, portalByCoords)

	portalsByLabel := make(map[string][]Portal)
	for _, portal := range portalByCoords {
		portalsByLabel[portal.label] = append(portalsByLabel[portal.label], portal)
	}

	// TODO/FIXME: should I do this?
	for coords := range portalByCoords {
		maze[coords.y][coords.x] = PortalDoor
	}

	return bfsFromTo(maze, portalByCoords, portalsByLabel, "AA", "ZZ")
}
