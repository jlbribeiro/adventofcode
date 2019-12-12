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

var bodyRegex = regexp.MustCompile(`<x=(\-?\d+), y=(\-?\d+), z=(\-?\d+)>`)

type Body struct {
	Pos [3]int
	Vel [3]int
}

func SimulateEnergyFromInput(input io.Reader, steps int) int {
	var bodies []*Body

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		bodyS := scanner.Text()

		posS := bodyRegex.FindStringSubmatch(bodyS)
		if len(posS) != 4 {
			panic(fmt.Errorf("unexpected input: %s", bodyS))
		}

		var pos [3]int
		for i, coordS := range posS[1:] {
			coord64, err := strconv.ParseInt(coordS, 10, 64)
			if err != nil {
				panic(err)
			}

			pos[i] = int(coord64)
		}

		bodies = append(bodies, &Body{
			Pos: pos,
			Vel: [3]int{0, 0, 0},
		})
	}

	return SimulateEnergy(bodies, steps)
}

func SimulateEnergy(bodies []*Body, steps int) int {
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
