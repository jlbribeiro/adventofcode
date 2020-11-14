package asciicamera

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/jlbribeiro/adventofcode/2019/day11/part2/thermal"
)

const scaffold rune = '#'
const maxCharacters int = 20

type VacuumRobot struct {
	x, y      int
	cpu       *thermal.CPU
	camera    [][]rune
	direction Direction
	path      []string
}

func (robot *VacuumRobot) boot() {
	photoA, _ := robot.cpu.Exec(nil)
	robot.camera = photoFromIntcodeOutput(photoA)

	robot.direction = North
	robot.x, robot.y = robot.selfLocate()
}

func (robot *VacuumRobot) selfLocate() (int, int) {
	for y := range robot.camera {
		for x := range robot.camera[y] {
			if robot.camera[y][x] == '^' {
				return x, y
			}
		}
	}

	return -1, -1
}

func (robot *VacuumRobot) nextDirection() (Direction, rune, bool) {
	turns := []rune{'L', 'R'}

	for _, turn := range turns {
		var direction Direction
		switch turn {
		case 'L':
			direction = robot.direction.Left()
		case 'R':
			direction = robot.direction.Right()
		}

		dy, dx := direction.Offsets()
		y, x := robot.y+dy, robot.x+dx
		if y < 0 || y >= len(robot.camera) || x < 0 || x >= len(robot.camera[y]) {
			continue
		}

		if robot.camera[y][x] == scaffold {
			return direction, turn, true
		}
	}

	return robot.direction, 'X', false
}

func (robot *VacuumRobot) walk() int {
	steps := 0
	dy, dx := robot.direction.Offsets()

	for {
		y, x := robot.y+dy, robot.x+dx
		if y < 0 || y >= len(robot.camera) || x < 0 || x >= len(robot.camera[y]) {
			return steps
		}

		if robot.camera[y][x] != scaffold {
			return steps
		}

		steps++
		robot.y, robot.x = y, x
	}
}

func isSubroutine(instruction string) bool {
	switch instruction {
	case "A", "B", "C":
		return true
	}

	return false
}

func instructionsFromPath(path []string) [][]rune {
	for instructionLen := maxCharacters; instructionLen >= 1; instructionLen-- {
		for begin := 0; begin < len(path); begin++ {
			if isSubroutine(path[begin]) {
				continue
			}

			curLen := len(path[begin])
			end := begin + 1
			for ; end < len(path); end++ {
				if isSubroutine(path[end]) {
					break
				}

				// "x x x" (curLen = 3) corresponds to 5 chars on the wire, "x, x, x"
				// (3 chars + 3 commas - the last comma)
				if (curLen+len(path[end]))+(end-begin) > instructionLen {
					break
				}

				curLen += len(path[end])
			}

			candidate := strings.Join(path[begin:end], ",")
			fmt.Printf("candidate: %s\n", candidate)
			fmt.Printf("begin:  %2d\tend: %2d\n", begin, end)
			fmt.Printf("curLen: %2d\tlen: %2d\n", curLen, len(candidate))
			fmt.Println()
		}
	}

	// // FIXME
	// instructions := [][]rune{
	// 	[]rune("A,A,B,C,B,C,B,C,C,A\n"),
	// 	[]rune("R,8,L,4,R,4,R,10,R,8\n"),
	// 	[]rune("L,12,L,12,R,8,R,8\n"),
	// 	[]rune("R,10,R,4,R,4\n"),
	// }

	// for _, instruction := range instructions {
	// 	instruction[len(instruction)-1] = '\n'
	// }

	// return instructions
	return [][]rune{}
}

func (robot *VacuumRobot) VacuumSpaceDust(out io.Writer) int {
	robot.boot()

	for {
		steps := robot.walk()
		if steps > 0 {
			robot.path = append(robot.path, strconv.Itoa(steps))
		}

		direction, turn, found := robot.nextDirection()
		if !found {
			break
		}

		robot.direction = direction
		robot.path = append(robot.path, string(turn))
	}

	fmt.Fprintln(out, strings.Join(robot.path, ","))

	instructions := instructionsFromPath(robot.path)

	var input []int64
	for _, instruction := range instructions {
		for _, c := range instruction {
			input = append(input, int64(c))
		}
	}

	// Continuous video feed?
	input = append(input, 'n')
	input = append(input, '\n')

	output, _ := robot.cpu.Exec(input)

	for _, c := range output[:len(output)-1] {
		fmt.Fprint(out, string(c))
	}

	return int(output[len(output)-1])
}

func photoFromIntcodeOutput(photoA []int64) [][]rune {
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

	return photo
}

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

	photoA, _ := cpu.Exec(nil)
	photo := photoFromIntcodeOutput(photoA)

	return SumAlignmentCoordinates(photo)
}

func VacuumSpaceDustFromInput(input io.Reader, out io.Writer) int {
	program := thermal.ProgramFromInput(input)
	program[0] = 2
	cpu := thermal.NewCPU(program)

	robot := &VacuumRobot{
		cpu: cpu,
	}
	return robot.VacuumSpaceDust(out)
}
