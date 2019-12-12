package nbody

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func gcd(a, b int64) int64 {
	for b != 0 {
		tmp := b
		b = a % b
		a = tmp
	}

	return a
}

func lcm(ints []int64) int64 {
	r := int64(1)

	for _, n := range ints {
		r = r * n / gcd(r, n)
	}

	return r
}

var bodyRegex = regexp.MustCompile(`<x=(\-?\d+), y=(\-?\d+), z=(\-?\d+)>`)

type Body struct {
	InitialPos [3]int
	Pos        [3]int
	Vel        [3]int
}

type System struct {
	Bodies []*Body
}

func NewSystemFromInput(input io.Reader) *System {
	var system System

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		bodyS := scanner.Text()

		posS := bodyRegex.FindStringSubmatch(bodyS)
		if len(posS) != 4 {
			panic(fmt.Errorf("unexpected input: %s", bodyS))
		}

		var pos [3]int
		for i, coordS := range posS[1:] {
			coord64, err := strconv.ParseInt(coordS, 10, 32)
			if err != nil {
				panic(err)
			}

			pos[i] = int(coord64)
		}

		system.Bodies = append(system.Bodies, &Body{
			InitialPos: pos,
			Pos:        pos,
			Vel:        [3]int{0, 0, 0},
		})
	}

	return &system
}

func (system *System) Simulate(steps int) {
	bodies := system.Bodies
	for step := 0; step < steps; step++ {
		for i := 0; i < len(bodies)-1; i++ {
			bodyA := bodies[i]
			for j := i + 1; j < len(bodies); j++ {
				bodyB := bodies[j]
				for axis := 0; axis < 3; axis++ {
					d := bodyA.Pos[axis] - bodyB.Pos[axis]
					if d == 0 {
						continue
					}

					d /= abs(d)
					bodyA.Vel[axis] += -d
					bodyB.Vel[axis] += d
				}
			}
		}

		for _, body := range bodies {
			for axis := 0; axis < 3; axis++ {
				body.Pos[axis] += body.Vel[axis]
			}
		}
	}
}

func (system *System) Energy() int {
	bodies := system.Bodies
	energy := 0
	for _, body := range bodies {
		potential := 0
		kinetic := 0

		for axis := 0; axis < 3; axis++ {
			potential += abs(body.Pos[axis])
			kinetic += abs(body.Vel[axis])
		}

		energy += potential * kinetic
	}

	return energy
}

func (system *System) SimulateUntilPreviousState() int64 {
	bodies := system.Bodies

	var cycles [3]int64
	foundCycles := 0

	for step := int64(1); ; step++ {
		system.Simulate(1)

		for axis := 0; axis < 3; axis++ {
			if cycles[axis] > 0 {
				continue
			}

			nBodiesSamePos := 0
			for _, body := range bodies {
				if body.Vel[axis] == 0 && body.Pos[axis] == body.InitialPos[axis] {
					nBodiesSamePos++
				}
			}

			if nBodiesSamePos == len(bodies) {
				cycles[axis] = step
				foundCycles++
			}
		}

		if foundCycles == 3 {
			break
		}
	}

	return lcm(cycles[:])
}

func SimulateEnergyFromInput(input io.Reader, steps int) int {
	system := NewSystemFromInput(input)
	system.Simulate(steps)
	return system.Energy()
}

func SimulateUntilPreviousStateFromInput(input io.Reader) int64 {
	system := NewSystemFromInput(input)
	return system.SimulateUntilPreviousState()
}
